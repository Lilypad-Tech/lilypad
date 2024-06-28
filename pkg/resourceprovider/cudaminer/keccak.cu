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

__device__ __forceinline__ uint64_t asm_cuda_keccak_ROTL64(const uint64_t x, const int offset) {
	uint64_t res;
	asm("{ // ROTL64 \n\t"
		".reg .u32 tl,th,vl,vh;\n\t"
		".reg .pred p;\n\t"
		"mov.b64 {tl,th}, %1;\n\t"
		"shf.l.wrap.b32 vl, tl, th, %2;\n\t"
		"shf.l.wrap.b32 vh, th, tl, %2;\n\t"
		"setp.lt.u32 p, %2, 32;\n\t"
		"@!p mov.b64 %0, {vl,vh};\n\t"
		"@p  mov.b64 %0, {vh,vl};\n\t"
	"}\n" : "=l"(res) : "l"(x) , "r"(offset)
	);
	return res;
}

__device__ __forceinline__ static void cuda_keccak_permutations(uint64_t* A)
{
    uint64_t *a00 = A, *a01 = A + 1, *a02 = A + 2, *a03 = A + 3, *a04 = A + 4;
    uint64_t *a05 = A + 5, *a06 = A + 6, *a07 = A + 7, *a08 = A + 8, *a09 = A + 9;
    uint64_t *a10 = A + 10, *a11 = A + 11, *a12 = A + 12, *a13 = A + 13, *a14 = A + 14;
    uint64_t *a15 = A + 15, *a16 = A + 16, *a17 = A + 17, *a18 = A + 18, *a19 = A + 19;
    uint64_t *a20 = A + 20, *a21 = A + 21, *a22 = A + 22, *a23 = A + 23, *a24 = A + 24;

    for (int i = 0; i < KECCAK_ROUND; i++) {

        /* Theta */
		uint64_t c0 = *a00^ *a05^ *a10^ *a15^ *a20;
		uint64_t c1 = *a01^ *a06^ *a11^ *a16^ *a21;
		uint64_t c2 = *a02^ *a07^ *a12^ *a17^ *a22;
		uint64_t c3 = *a03^ *a08^ *a13^ *a18^ *a23;
		uint64_t c4 =*a04^ *a09^ *a14^ *a19^ *a24;
		
        int64_t d1 = asm_cuda_keccak_ROTL64(c1, 1) ^ c4;
        int64_t d2 = asm_cuda_keccak_ROTL64(c2, 1) ^ c0;
        int64_t d3 = asm_cuda_keccak_ROTL64(c3, 1) ^ c1;
        int64_t d4 = asm_cuda_keccak_ROTL64(c4, 1) ^ c2;
        int64_t d0 = asm_cuda_keccak_ROTL64(c0, 1) ^ c3;

        *a00 ^= d1;
        *a05 ^= d1;
        *a10 ^= d1;
        *a15 ^= d1;
        *a20 ^= d1;
        *a01 ^= d2;
        *a06 ^= d2;
        *a11 ^= d2;
        *a16 ^= d2;
        *a21 ^= d2;
        *a02 ^= d3;
        *a07 ^= d3;
        *a12 ^= d3;
        *a17 ^= d3;
        *a22 ^= d3;
        *a03 ^= d4;
        *a08 ^= d4;
        *a13 ^= d4;
        *a18 ^= d4;
        *a23 ^= d4;
        *a04 ^= d0;
        *a09 ^= d0;
        *a14 ^= d0;
        *a19 ^= d0;
        *a24 ^= d0;


        /* Rho pi */
        c1 = asm_cuda_keccak_ROTL64(*a01, 1);
        *a01 = asm_cuda_keccak_ROTL64(*a06, 44);
        *a06 = asm_cuda_keccak_ROTL64(*a09, 20);
        *a09 = asm_cuda_keccak_ROTL64(*a22, 61);
        *a22 = asm_cuda_keccak_ROTL64(*a14, 39);
        *a14 = asm_cuda_keccak_ROTL64(*a20, 18);
        *a20 = asm_cuda_keccak_ROTL64(*a02, 62);
        *a02 = asm_cuda_keccak_ROTL64(*a12, 43);
        *a12 = asm_cuda_keccak_ROTL64(*a13, 25);
        *a13 = asm_cuda_keccak_ROTL64(*a19, 8);
        *a19 = asm_cuda_keccak_ROTL64(*a23, 56);
        *a23 = asm_cuda_keccak_ROTL64(*a15, 41);
        *a15 = asm_cuda_keccak_ROTL64(*a04, 27);
        *a04 = asm_cuda_keccak_ROTL64(*a24, 14);
        *a24 = asm_cuda_keccak_ROTL64(*a21, 2);
        *a21 = asm_cuda_keccak_ROTL64(*a08, 55);
        *a08 = asm_cuda_keccak_ROTL64(*a16, 45);
        *a16 = asm_cuda_keccak_ROTL64(*a05, 36);
        *a05 = asm_cuda_keccak_ROTL64(*a03, 28);
        *a03 = asm_cuda_keccak_ROTL64(*a18, 21);
        *a18 = asm_cuda_keccak_ROTL64(*a17, 15);
        *a17 = asm_cuda_keccak_ROTL64(*a11, 10);
        *a11 = asm_cuda_keccak_ROTL64(*a07, 6);
        *a07 = asm_cuda_keccak_ROTL64(*a10, 3);
        *a10 = c1;

        /* Chi * a ^ (~b) & c*/  
        c0 = *a00 ^ (~*a01 & *a02);  // use int2 vector this can be opt to 2 lop.b32 instruction
        c1 = *a01 ^ (~*a02 & *a03);
        *a02 ^= ~*a03 & *a04;
        *a03 ^= ~*a04 & *a00;
        *a04 ^= ~*a00 & *a01;
        *a00 = c0;
        *a01 = c1;

        c0 = *a05 ^ (~*a06 & *a07);
        c1 = *a06 ^ (~*a07 & *a08);
        *a07 ^= ~*a08 & *a09;
        *a08 ^= ~*a09 & *a05;
        *a09 ^= ~*a05 & *a06;
        *a05 = c0;
        *a06 = c1;

        c0 = *a10 ^ (~*a11 & *a12);
        c1 = *a11 ^ (~*a12 & *a13);
        *a12 ^= ~*a13 & *a14;
        *a13 ^= ~*a14 & *a10;
        *a14 ^= ~*a10 & *a11;
        *a10 = c0;
        *a11 = c1;

        c0 = *a15 ^ (~*a16 & *a17);
        c1 = *a16 ^ (~*a17 & *a18);
        *a17 ^= ~*a18 & *a19;
        *a18 ^= ~*a19 & *a15;
        *a19 ^= ~*a15 & *a16;
        *a15 = c0;
        *a16 = c1;

        c0 = *a20 ^ (~*a21 & *a22);
        c1 = *a21 ^ (~*a22 & *a23);
        *a22 ^= ~*a23 & *a24;
        *a23 ^= ~*a24 & *a20;
        *a24 ^= ~*a20 & *a21;
        *a20 = c0;
        *a21 = c1;

        /* Iota */
        *a00 ^= CUDA_KECCAK_CONSTS[i];
    }
}

__noinline__ __device__ static bool hashbelowtarget(const uint64_t *const __restrict__ hash, const uint64_t *const __restrict__ target)
{
    if (hash[3] > target[3])
        return false;
    if (hash[3] < target[3])
        return true;
    if (hash[2] > target[2])
        return false;
    if (hash[2] < target[2])
        return true;

    if (hash[1] > target[1])
        return false;
    if (hash[1] < target[1])
        return true;
    if (hash[0] > target[0])
        return false;

    return true;
}

__device__ uint64_t *addUint256(const uint64_t *a, const uint64_t b)
{
    uint64_t *result = new uint64_t[4];
    uint64_t sum = a[0] + b;
    result[0] = sum;

    uint64_t carry = (sum < a[0]) ? 1 : 0;
    for (int i = 1; i < 4; i++)
    {
        sum = a[i] + carry;
        result[i] = sum;
        carry = (sum < a[i]) ? 1 : 0;
    }

    return result;
}

__device__ void reverseArray(unsigned char *array, int n) {
    for (int i = 0; i < n / 2; ++i) {
        unsigned char temp = array[i];
        array[i] = array[n - 1 - i];
        array[n - 1 - i] = temp;
    }
}


extern "C" __global__ __launch_bounds__(1024, 1)
  void kernel_lilypad_pow(uint8_t* challenge, uint64_t* startNonce,  uint64_t* target, uint32_t n_batch, uint8_t* resNonce)
{
    uint32_t thread = blockIdx.x * blockDim.x + threadIdx.x; 
    if (thread >= n_batch) {
        return;
    }

       //increase nonce
    uint8_t* nonce = (uint8_t*)addUint256(startNonce, thread);
    uint64_t state[KECCAK_STATE_SIZE];
    memset(state, 0, sizeof(state));

    memcpy(state, challenge, 32);  // Copy challenge into state
    memcpy(state + 4, nonce, 32);  // Copy nonce into state starting from index 4

    state[8] ^= 1;
    state[16] ^= 9223372036854775808ULL; 

    cuda_keccak_permutations(state);

    uint8_t out[32];
    uint8_t* state_bytes = reinterpret_cast<uint8_t*>(state);
    #pragma unroll 32
    for (int i = 0;i<32; i++) {
        out[i] = state_bytes[31-i];
    }
    
    if (hashbelowtarget((uint64_t*)out, target)) {
        memcpy(resNonce, nonce, 32);
    } 

    delete nonce;//45
}


extern "C" __global__ __launch_bounds__(1024, 1)
  void kernel_lilypad_pow_debug(uint8_t* challenge, uint64_t* startNonce,  uint64_t* target, uint32_t n_batch, uint8_t* resNonce,  uint8_t *hash, uint8_t *pack)
{
    uint32_t thread = blockIdx.x * blockDim.x + threadIdx.x; 
    if (thread >= n_batch) {
        return;
    }

       //increase nonce
    uint8_t* nonce = (uint8_t*)addUint256(startNonce, thread);
    uint64_t state[KECCAK_STATE_SIZE];
    memset(state, 0, sizeof(state));

    memcpy(state, challenge, 32);  // Copy challenge into state
    memcpy(state + 4, nonce, 32);  // Copy nonce into state starting from index 4

    //uint8_t cuda_pack[64];
    //memcpy(cuda_pack, state, 64);

    state[8] ^= 1;
    state[16] ^= 9223372036854775808ULL; 

    cuda_keccak_permutations(state);

    uint8_t out[32];
    uint8_t* state_bytes = reinterpret_cast<uint8_t*>(state);
    #pragma unroll 32
    for (int i = 0;i<32; i++) {
        out[i] = state_bytes[31-i];
    }
    
    if (hashbelowtarget((uint64_t*)out, target)) {
       // reverseArray(out, 32);
       // memcpy(hash, out, 32);
       // memcpy(pack, cuda_pack, 64);
        memcpy(resNonce, nonce, 32);
    } 

    delete nonce;//45
}
