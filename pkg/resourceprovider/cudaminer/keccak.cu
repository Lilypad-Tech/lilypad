/*
 * keccak.cu  Implementation of Keccak/SHA3 digest
 *
 * Date: 12 June 2019
 * Revision: 1
 *
 * This file is released into the Public Domain.
 */
 
 
extern "C"
{
    #include "keccak.cuh"
}

#define KECCAK_ROUND 24
#define KECCAK_STATE_SIZE 25
#define KECCAK_Q_SIZE 192

__constant__ uint64_t CUDA_KECCAK_CONSTS[24] = { 0x0000000000000001, 0x0000000000008082,
                                          0x800000000000808a, 0x8000000080008000, 0x000000000000808b, 0x0000000080000001, 0x8000000080008081,
                                          0x8000000000008009, 0x000000000000008a, 0x0000000000000088, 0x0000000080008009, 0x000000008000000a,
                                          0x000000008000808b, 0x800000000000008b, 0x8000000000008089, 0x8000000000008003, 0x8000000000008002,
                                          0x8000000000000080, 0x000000000000800a, 0x800000008000000a, 0x8000000080008081, 0x8000000000008080,
                                          0x0000000080000001, 0x8000000080008008 };


__constant__  uint64_t digestbitlen = 256;
__constant__  uint64_t rate_bits = 1088;
__constant__  uint64_t rate_BYTEs = 136;
__constant__  uint64_t absorb_round = 17;
typedef struct {
    int64_t state[KECCAK_STATE_SIZE];
    uint8_t q[KECCAK_Q_SIZE];

    uint64_t bits_in_queue;

} cuda_keccak_ctx_t;
typedef cuda_keccak_ctx_t CUDA_KECCAK_CTX;

__device__ uint64_t cuda_keccak_leuint64(void *in)
{
    uint64_t a;
    memcpy(&a, in, 8);
    return a;
}

__device__ int64_t cuda_keccak_MIN(int64_t a, int64_t b)
{
    if (a > b) return b;
    return a;
}

__device__ uint64_t cuda_keccak_UMIN(uint64_t a, uint64_t b)
{
    if (a > b) return b;
    return a;
}

__device__ void cuda_keccak_extract(cuda_keccak_ctx_t *ctx)
{
    uint64_t len = rate_bits >> 6;
    int64_t a;
    int s = sizeof(uint64_t);

    for (int i = 0;i < len;i++) {
        a = cuda_keccak_leuint64((int64_t*)&ctx->state[i]);
        memcpy(ctx->q + (i * s), &a, s);
    }
}
__device__ __forceinline__ uint64_t cuda_keccak_ROTL64(uint64_t a, uint64_t b) {
    return (a << b) | (a >> (64 - b));
}

__device__ void cuda_keccak_permutations(cuda_keccak_ctx_t *ctx) {
    int64_t* A = ctx->state;

    #pragma unroll 24
    for (int i = 0; i < KECCAK_ROUND; i++) {
        int64_t C[5], D[5];  

        // Theta
        C[0] = A[0] ^ A[5] ^ A[10] ^ A[15] ^ A[20];
        C[1] = A[1] ^ A[6] ^ A[11] ^ A[16] ^ A[21];
        C[2] = A[2] ^ A[7] ^ A[12] ^ A[17] ^ A[22];
        C[3] = A[3] ^ A[8] ^ A[13] ^ A[18] ^ A[23];
        C[4] = A[4] ^ A[9] ^ A[14] ^ A[19] ^ A[24];

        D[0] = cuda_keccak_ROTL64(C[1], 1) ^ C[4];
        D[1] = cuda_keccak_ROTL64(C[2], 1) ^ C[0];
        D[2] = cuda_keccak_ROTL64(C[3], 1) ^ C[1];
        D[3] = cuda_keccak_ROTL64(C[4], 1) ^ C[2];
        D[4] = cuda_keccak_ROTL64(C[0], 1) ^ C[3];

        #pragma unroll 25
        for (int j = 0; j < 25; j += 5) {
            A[j] ^= D[0];
            A[j + 1] ^= D[1];
            A[j + 2] ^= D[2];
            A[j + 3] ^= D[3];
            A[j + 4] ^= D[4];
        }

        // Rho Pi
        int64_t B[25];
        B[0] = A[0];
        B[1] = cuda_keccak_ROTL64(A[6], 44);
        B[2] = cuda_keccak_ROTL64(A[12], 43);
        B[3] = cuda_keccak_ROTL64(A[18], 21);
        B[4] = cuda_keccak_ROTL64(A[24], 14);
        B[5] = cuda_keccak_ROTL64(A[3], 28);
        B[6] = cuda_keccak_ROTL64(A[9], 20);
        B[7] = cuda_keccak_ROTL64(A[10], 3);
        B[8] = cuda_keccak_ROTL64(A[16], 45);
        B[9] = cuda_keccak_ROTL64(A[22], 61);
        B[10] = cuda_keccak_ROTL64(A[1], 1);
        B[11] = cuda_keccak_ROTL64(A[7], 6);
        B[12] = cuda_keccak_ROTL64(A[13], 25);
        B[13] = cuda_keccak_ROTL64(A[19], 8);
        B[14] = cuda_keccak_ROTL64(A[20], 18);
        B[15] = cuda_keccak_ROTL64(A[4], 27);
        B[16] = cuda_keccak_ROTL64(A[5], 36);
        B[17] = cuda_keccak_ROTL64(A[11], 10);
        B[18] = cuda_keccak_ROTL64(A[17], 15);
        B[19] = cuda_keccak_ROTL64(A[23], 56);
        B[20] = cuda_keccak_ROTL64(A[2], 62);
        B[21] = cuda_keccak_ROTL64(A[8], 55);
        B[22] = cuda_keccak_ROTL64(A[14], 39);
        B[23] = cuda_keccak_ROTL64(A[15], 41);
        B[24] = cuda_keccak_ROTL64(A[21], 2);

        // Chi
        #pragma unroll 24
        for (int j = 0; j < 25; j += 5) {
            #pragma unroll 5
            for (int k = 0; k < 5; ++k) {
                A[j + k] = B[j + k] ^ (~B[j + (k + 1) % 5] & B[j + (k + 2) % 5]);
            }
        }

        // Iota
        A[0] ^= CUDA_KECCAK_CONSTS[i];
    }
}


__device__ void cuda_keccak_absorb(cuda_keccak_ctx_t *ctx, uint8_t* in)
{

    uint64_t offset = 0;
    for (uint64_t i = 0; i < absorb_round; ++i) {//10
        ctx->state[i] ^= cuda_keccak_leuint64(in + offset);//18
        offset += 8;//9
    }

    cuda_keccak_permutations(ctx);//8
}

__device__ void cuda_keccak_pad(cuda_keccak_ctx_t *ctx)
{
    ctx->q[ctx->bits_in_queue >> 3] |= (1L << (ctx->bits_in_queue & 7)); //6

    if (++(ctx->bits_in_queue) == rate_bits) {//9
        cuda_keccak_absorb(ctx, ctx->q);//8
        ctx->bits_in_queue = 0;//53
    }

    uint64_t full = ctx->bits_in_queue >> 6;    //7
    uint64_t partial = ctx->bits_in_queue & 63; //8

    uint64_t offset = 0;
    for (int i = 0; i < full; ++i) {//52
        ctx->state[i] ^= cuda_keccak_leuint64(ctx->q + offset);//52
        offset += 8;//52
    }

    if (partial > 0) {//8
        uint64_t mask = (1L << partial) - 1;//17
        ctx->state[full] ^= cuda_keccak_leuint64(ctx->q + offset) & mask;//16
    }

    ctx->state[(rate_bits - 1) >> 6] ^= 9223372036854775808ULL;/* 1 << 63 */   //9

    cuda_keccak_permutations(ctx);//8
    cuda_keccak_extract(ctx);//58

    ctx->bits_in_queue = rate_bits;//37
}


/*
 * Digestbitlen must be 128 224 256 288 384 512
 */
__device__ void cuda_keccak_init(cuda_keccak_ctx_t *ctx)
{
    memset(ctx, 0, sizeof(cuda_keccak_ctx_t));
    ctx->bits_in_queue = 0;//11
}

__device__ void cuda_keccak_update(cuda_keccak_ctx_t *ctx, uint8_t *in, uint64_t inlen)
{
    int64_t BYTEs = ctx->bits_in_queue >> 3;
    int64_t count = 0;
    while (count < inlen) {//46
        if (BYTEs == 0 && count <= ((int64_t)(inlen - rate_BYTEs))) {//12
            do {
                cuda_keccak_absorb(ctx, in + count);//8
                count += rate_BYTEs;//56
            } while (count <= ((int64_t)(inlen - rate_BYTEs)));//46
        } else {
            int64_t partial = cuda_keccak_MIN(rate_BYTEs - BYTEs, inlen - count);//12
            memcpy(ctx->q + BYTEs, in + count, partial);//12

            BYTEs += partial;//10
            count += partial;//8

            if (BYTEs == rate_BYTEs) {//10
                cuda_keccak_absorb(ctx, ctx->q);//8
                BYTEs = 0;
            }
        }
    }
    ctx->bits_in_queue = BYTEs << 3;//8
}

__device__ void cuda_keccak_final_rev(cuda_keccak_ctx_t *ctx, uint8_t *out)
{
    cuda_keccak_pad(ctx);
    uint64_t i = 0;//6

    while (i < digestbitlen) {//46
        if (ctx->bits_in_queue == 0) {//9
            cuda_keccak_permutations(ctx);//8
            cuda_keccak_extract(ctx);//56
            ctx->bits_in_queue = rate_bits;//7
        }

        uint64_t partial_block = cuda_keccak_UMIN(ctx->bits_in_queue, digestbitlen - i);//9

        //directly reverse?
        int start = 31- (i >> 3);
        uint8_t* pos = ctx->q + (rate_BYTEs - (ctx->bits_in_queue >> 3));
        for (int j = 0; j< (partial_block >> 3); j++) {
            out[start-j] = pos[j];
        }

        ctx->bits_in_queue -= partial_block;//11
        i += partial_block;//11
    }
}



__noinline__ __device__ static bool hashbelowtarget(const uint64_t *const __restrict__ hash, const uint64_t *const __restrict__ target)
{
    if (hash[3] > target[3])//46
        return false;
    if (hash[3] < target[3])//46
        return true;
    if (hash[2] > target[2])//45
        return false;
    if (hash[2] < target[2])//45
        return true;

    if (hash[1] > target[1])//43
        return false;
    if (hash[1] < target[1])//43
        return true;
    if (hash[0] > target[0])//39
        return false;

    return true;
}

__device__ uint64_t *addUint256(const uint64_t *a, const uint64_t b)
{
    uint64_t *result = new uint64_t[4];//47
    uint64_t sum = a[0] + b;//10
    result[0] = sum;//10

    uint64_t carry = (sum < a[0]) ? 1 : 0;//12
    for (int i = 1; i < 4; i++)//13
    {
        sum = a[i] + carry;//16
        result[i] = sum;//14
        carry = (sum < a[i]) ? 1 : 0;//14
    }

    return result;
}
__device__ void reverse32BytesInPlace(uint8_t *data, uint8_t *out)
{
    for (int i = 0; i < 32; i++)//13
    {
       out[i] = data[31-i];
    }
}

extern "C" __global__ __launch_bounds__(1024)

  void kernel_lilypad_pow(uint8_t* chanllenge, uint64_t* startNonce,  uint64_t* target, uint64_t n_batch, uint8_t* resNonce)
{
    uint64_t thread = blockIdx.x * blockDim.x + threadIdx.x; //4
    if (thread >= n_batch) {//36
        return;
    }

    //pack input
    uint8_t in[64];
    memcpy(in, chanllenge, 32);
    //increase nonce
    uint8_t* nonce = (uint8_t*)addUint256(startNonce, thread);//35
    uint8_t nonce_rev[32];
    reverse32BytesInPlace(nonce, nonce_rev);//18
    memcpy(in+32, nonce_rev, 32);
    

    uint8_t out[32];
    CUDA_KECCAK_CTX ctx;
    cuda_keccak_init(&ctx);        //6
    cuda_keccak_update(&ctx, in,64);   //12
    cuda_keccak_final_rev(&ctx, out);       //6

    if (hashbelowtarget((uint64_t*)out, target)) {//49
        memcpy(resNonce, nonce_rev, 32);
    } 

    delete nonce;//45
}


extern "C" __global__ __launch_bounds__(1024)

  void kernel_lilypad_pow_debug(uint8_t* chanllenge, uint64_t* startNonce,  uint64_t* target, uint64_t n_batch, uint8_t* resNonce,  uint8_t *hash, uint8_t *pack)
{
    uint64_t thread = blockIdx.x * blockDim.x + threadIdx.x; //4
    if (thread >= n_batch) {//36
        return;
    }

    //pack input
    uint8_t in[64];
    memcpy(in, chanllenge, 32);
    //increase nonce
    uint8_t* nonce = (uint8_t*)addUint256(startNonce, thread);//35
    uint8_t nonce_rev[32];
    reverse32BytesInPlace(nonce, nonce_rev);//18
    memcpy(in+32, nonce_rev, 32);
    

    uint8_t out[32];
    CUDA_KECCAK_CTX ctx;
    cuda_keccak_init(&ctx);        //6
    cuda_keccak_update(&ctx, in,64);   //12
    cuda_keccak_final_rev(&ctx, out);       //6

    if (hashbelowtarget((uint64_t*)out, target)) {//49
      //  uint8_t out_rev[64];
      //  reverse32BytesInPlace(out, out_rev);//18
       // memcpy(hash, out_rev, 32);
      //  memcpy(pack, in, 64);
        memcpy(resNonce, nonce_rev, 32);
    } 

    delete nonce;//45
}
