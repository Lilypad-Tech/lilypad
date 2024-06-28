/*
 * Type Definitions for CUDA Hashing Algos
 *
 * Date: 12 June 2019
 * Revision: 1
 *
 * This file is released into the Public Domain.
 */

#pragma once
#define USE_MD2 1
#define USE_MD5 1
#define USE_SHA1 1
#define USE_SHA256 1

#define CUDA_HASH 1
#define OCL_HASH 0

typedef unsigned char BYTE;
typedef unsigned int  WORD;
typedef unsigned long long LONG;

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdint.h>