package bitboard

// Board is representing 8x8 (64 bit) bitboard where each bit represent existing piece on the given position
type Board uint64

// BitWidth holds number of bits in the bitboard
const BitWidth = 64
