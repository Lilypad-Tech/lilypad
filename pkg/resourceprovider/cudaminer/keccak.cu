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

typedef union
{
    uint2 uint2;
    uint64_t uint64;
    uint8_t uint8[8];
} nonce_t;

__constant__ uint64_t d_midstate[25];
__constant__ uint64_t d_target[1];

__device__ __forceinline__ nonce_t bswap_64(nonce_t const input)
{
    nonce_t output;
    asm("{"
        "  prmt.b32 %0, %3, 0, 0x0123;"
        "  prmt.b32 %1, %2, 0, 0x0123;"
        "}" : "=r"(output.uint2.x), "=r"(output.uint2.y) : "r"(input.uint2.x), "r"(input.uint2.y));
    return output;
}

__device__ __forceinline__ nonce_t xor5(nonce_t const a, nonce_t const b, nonce_t const c, nonce_t const d, nonce_t const e)
{
    nonce_t output;
    asm("{"
        "  lop3.b32 %0, %2, %4, %6, 0x96;"
        "  lop3.b32 %1, %3, %5, %7, 0x96;"
        "  lop3.b32 %0, %0, %8, %10, 0x96;"
        "  lop3.b32 %1, %1, %9, %11, 0x96;"
        "}" : "=r"(output.uint2.x), "=r"(output.uint2.y)
        : "r"(a.uint2.x), "r"(a.uint2.y), "r"(b.uint2.x), "r"(b.uint2.y), "r"(c.uint2.x), "r"(c.uint2.y), "r"(d.uint2.x), "r"(d.uint2.y), "r"(e.uint2.x), "r"(e.uint2.y));
    return output;
}

__device__ __forceinline__ nonce_t xor3(nonce_t const a, nonce_t const b, nonce_t const c)
{
    nonce_t output;
    asm("{"
        "  lop3.b32 %0, %2, %4, %6, 0x96;"
        "  lop3.b32 %1, %3, %5, %7, 0x96;"
        "}" : "=r"(output.uint2.x), "=r"(output.uint2.y)
        : "r"(a.uint2.x), "r"(a.uint2.y), "r"(b.uint2.x), "r"(b.uint2.y), "r"(c.uint2.x), "r"(c.uint2.y));
    return output;
}

__device__ __forceinline__ nonce_t chi(nonce_t const a, nonce_t const b, nonce_t const c)
{
    nonce_t output;
    asm("{"
        "  lop3.b32 %0, %2, %4, %6, 0xD2;"
        "  lop3.b32 %1, %3, %5, %7, 0xD2;"
        "}" : "=r"(output.uint2.x), "=r"(output.uint2.y)
        : "r"(a.uint2.x), "r"(a.uint2.y), "r"(b.uint2.x), "r"(b.uint2.y), "r"(c.uint2.x), "r"(c.uint2.y));

    return output;
}

__device__ __forceinline__ nonce_t rotl(nonce_t input, uint32_t const offset)
{
    asm("{"
        "  .reg .b32 tmp;"
        "  shf.l.wrap.b32 tmp, %1, %0, %2;"
        "  shf.l.wrap.b32 %1, %0, %1, %2;"
        "  mov.b32 %0, tmp;"
        "}" : "+r"(input.uint2.x), "+r"(input.uint2.y) : "r"(offset));
    return input;
}

__device__ __forceinline__ nonce_t rotr(nonce_t input, uint32_t const offset)
{
    asm("{"
        "  .reg .b32 tmp;"
        "  shf.r.wrap.b32 tmp, %0, %1, %2;"
        "  shf.r.wrap.b32 %1, %1, %0, %2;"
        "  mov.b32 %0, tmp;"
        "}" : "+r"(input.uint2.x), "+r"(input.uint2.y) : "r"(offset));
    return input;
}

__device__ uint64_t rotate(uint64_t val, unsigned n) { return val << n | val >> (64 - n); }

// Array of indices and rotation values for P and Pi phases.
__constant__ uint8_t g_ppi_aux[25][2] = {
    {0, 0}, {6, 44}, {12, 43}, {18, 21}, {24, 14}, {3, 28}, {9, 20}, {10, 3}, {16, 45}, {22, 61}, {1, 1}, {7, 6}, {13, 25}, {19, 8}, {20, 18}, {4, 27}, {5, 36}, {11, 10}, {17, 15}, {23, 56}, {2, 62}, {8, 55}, {14, 39}, {15, 41}, {21, 2}};

// Array of indices for ksi phase.
__constant__ uint8_t g_ksi_aux[25][2] = {
    {1, 2}, {2, 3}, {3, 4}, {4, 0}, {0, 1}, {6, 7}, {7, 8}, {8, 9}, {9, 5}, {5, 6}, {11, 12}, {12, 13}, {13, 14}, {14, 10}, {10, 11}, {16, 17}, {17, 18}, {18, 19}, {19, 15}, {15, 16}, {21, 22}, {22, 23}, {23, 24}, {24, 20}, {20, 21}};

__constant__ uint64_t g_iota_aux[24] = {
    0x0000000000000001L, 0x0000000000008082L, 0x800000000000808aL, 0x8000000080008000L, 0x000000000000808bL,
    0x0000000080000001L, 0x8000000080008081L, 0x8000000000008009L, 0x000000000000008aL, 0x0000000000000088L,
    0x0000000080008009L, 0x000000008000000aL, 0x000000008000808bL, 0x800000000000008bL, 0x8000000000008089L,
    0x8000000000008003L, 0x8000000000008002L, 0x8000000000000080L, 0x000000000000800aL, 0x800000008000000aL,
    0x8000000080008081L, 0x8000000000008080L, 0x0000000080000001L, 0x8000000080008008L};

__device__ static void cuda_keccak_permutations(nonce_t *A, nonce_t *C, const int threadIndexInWrap)
{
    size_t s = threadIndexInWrap % 5;
#pragma unroll
    for (int round_idx = 0; round_idx < 24; ++round_idx)
    {
        // Thetta phase.
        C[threadIndexInWrap] = xor5(A[s], A[s + 5], A[s + 10], A[s + 15], A[s + 20]);
        A[threadIndexInWrap] = xor3(A[threadIndexInWrap], C[s + 5 - 1], rotl(C[s + 1], 1));

        // P and Pi combined phases.
        C[threadIndexInWrap].uint64 = rotate(A[g_ppi_aux[threadIndexInWrap][0]].uint64, g_ppi_aux[threadIndexInWrap][1]);

        // Ksi phase.
        A[threadIndexInWrap] = chi(C[threadIndexInWrap], C[g_ksi_aux[threadIndexInWrap][0]], C[g_ksi_aux[threadIndexInWrap][1]]);

        // Iota phase.
        if (threadIndexInWrap == 0)
        {
            A[threadIndexInWrap].uint64 ^= g_iota_aux[round_idx];
        }
    }
}

__device__ static bool hashbelowtarget(const uint8_t *const __restrict__ hash, const uint8_t *const __restrict__ target)
{
    for (int i = 0; i < 32; i++)
    {
        if (hash[i] < target[i])
        {
            return true;
        }
        else if (hash[i] > target[i])
        {
            return false;
        }
    }
    return false;
}

__device__ __noinline__ void addUint256(nonce_t *result, const uint64_t *a, uint64_t b)
{
    uint64_t sum = a[0] + b;
    result[0].uint64 = sum;

    uint64_t carry = (sum < a[0]) ? 1 : 0;
    for (int i = 1; i < 4; i++)
    {
        sum = a[i] + carry;
        result[i].uint64 = sum;
        carry = (sum < a[i]) ? 1 : 0;
    }
}


#define WRAP_IN_BLOCK 32 // equal to block_size/32

extern "C" __global__ __launch_bounds__(1024) void kernel_lilypad_pow(
    const uint8_t *__restrict__ challenge,
    const uint64_t *__restrict__ startNonce,
    const uint8_t *__restrict__ target,
    const uint32_t n_batch,
    const uint32_t hashPerThread, uint8_t *resNonce)
{
    int thread = blockIdx.x * blockDim.x + threadIdx.x;
    if (thread >= n_batch) // batch must equal with grid*block
    {
        return;
    }

    uint64_t wrapInOneLaunch = thread / 32;
    int threadIndexInWrap = thread % 32; // index in wrap
    if (threadIndexInWrap >= 25)         // abort 26-32 thread
    {
        return;
    }

    int wrapIndexInBlock = threadIdx.x / 32; // one wrap one worker, 25/32 usages

    __shared__ nonce_t stateInBlock[WRAP_IN_BLOCK][KECCAK_STATE_SIZE];
    __shared__ nonce_t cInBlock[WRAP_IN_BLOCK][25];

    nonce_t *state = stateInBlock[wrapIndexInBlock];
    nonce_t *C = cInBlock[wrapIndexInBlock];

    C[threadIndexInWrap].uint64 = 0;

    __syncwarp();
    uint64_t nonceOffset = wrapInOneLaunch * hashPerThread;
    uint64_t endNonceOffset = (wrapInOneLaunch + 1) * hashPerThread;
    for (; nonceOffset < endNonceOffset; nonceOffset++)
    {
        nonce_t nonce[4];

        state[threadIndexInWrap].uint64 = 0;
        if (threadIndexInWrap == 0)
        {
            // increase nonce
            addUint256(nonce, startNonce, nonceOffset);
            memcpy(state, challenge, 32); // Copy challenge into state
            memcpy(state + 4, nonce, 32); // Copy nonce into state starting from index 4

            state[8].uint64 ^= 1;
            state[16].uint64 ^= 9223372036854775808ULL;
        }

        __syncwarp();
        cuda_keccak_permutations(state, C, threadIndexInWrap);

        if (threadIndexInWrap == 0)
        {

            if (hashbelowtarget(state->uint8, target))
            {
                memcpy(resNonce, nonce, 32);
            }

            delete nonce; // 45
        }
    }
}

extern "C" __global__ __launch_bounds__(1024) void kernel_lilypad_pow_debug(
    const uint8_t *__restrict__ challenge,
    const uint64_t *__restrict__ startNonce,
    const uint8_t *__restrict__ target,
    const uint32_t n_batch,
    const uint32_t hashPerThread, uint8_t *resNonce, uint8_t *hash, uint8_t *pack)
{
    int thread = blockIdx.x * blockDim.x + threadIdx.x;
    if (thread >= n_batch) // batch must equal with grid*block
    {
        return;
    }

    uint64_t wrapInOneLaunch = thread / 32;
    int threadIndexInWrap = thread % 32; // index in wrap
    if (threadIndexInWrap >= 25)         // abort 26-32 thread
    {
        return;
    }

    int wrapIndexInBlock = threadIdx.x / 32; // one wrap one worker, 25/32 usages

    __shared__ nonce_t stateInBlock[WRAP_IN_BLOCK][KECCAK_STATE_SIZE];
    __shared__ nonce_t cInBlock[WRAP_IN_BLOCK][25];

    nonce_t *state = stateInBlock[wrapIndexInBlock];
    nonce_t *C = cInBlock[wrapIndexInBlock];

    C[threadIndexInWrap].uint64 = 0;

    __syncwarp();
    uint64_t nonceOffset = wrapInOneLaunch * hashPerThread;
    uint64_t endNonceOffset = (wrapInOneLaunch + 1) * hashPerThread;
    for (; nonceOffset < endNonceOffset; nonceOffset++)
    {
        uint8_t cuda_pack[64];
        nonce_t nonce[4];

        state[threadIndexInWrap].uint64 = 0;
        if (threadIndexInWrap == 0)
        {
            // increase nonce
            addUint256(nonce, startNonce, nonceOffset);
            memcpy(state, challenge, 32); // Copy challenge into state
            memcpy(state + 4, nonce, 32); // Copy nonce into state starting from index 4

            memcpy(cuda_pack, state, 64);

            state[8].uint64 ^= 1;
            state[16].uint64 ^= 9223372036854775808ULL;
        }

        __syncwarp();
        cuda_keccak_permutations(state, C, threadIndexInWrap);

        if (threadIndexInWrap == 0)
        {

            if (hashbelowtarget(state->uint8, target))
            {
                memcpy(hash, state, 32);
                memcpy(pack, cuda_pack, 64);
                memcpy(resNonce, nonce, 32);
            }

            delete nonce; // 45
        }
    }
}
