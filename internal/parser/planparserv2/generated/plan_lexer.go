// Code generated from Plan.g4 by ANTLR 4.9. DO NOT EDIT.

package planparserv2

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 42, 530,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 160, 10, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	5, 14, 174, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 5, 26, 206, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 212,
	10, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 220, 10, 29, 3,
	30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 32, 7, 32, 235, 10, 32, 12, 32, 14, 32, 238, 11, 32, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 268, 10, 33, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 297, 10, 34, 3, 35, 3, 35, 3, 35, 3,
	35, 5, 35, 303, 10, 35, 3, 36, 3, 36, 5, 36, 307, 10, 36, 3, 37, 3, 37,
	3, 37, 7, 37, 312, 10, 37, 12, 37, 14, 37, 315, 11, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 5, 37, 322, 10, 37, 3, 38, 5, 38, 325, 10, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 334, 10, 38, 3, 38, 3,
	38, 7, 38, 338, 10, 38, 12, 38, 14, 38, 341, 11, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 352, 10, 38, 3, 38, 3,
	38, 3, 38, 3, 38, 7, 38, 358, 10, 38, 12, 38, 14, 38, 361, 11, 38, 3, 38,
	3, 38, 5, 38, 365, 10, 38, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 371, 10,
	39, 3, 39, 3, 39, 6, 39, 375, 10, 39, 13, 39, 14, 39, 376, 3, 40, 3, 40,
	3, 40, 5, 40, 382, 10, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3,
	43, 6, 43, 391, 10, 43, 13, 43, 14, 43, 392, 3, 44, 3, 44, 7, 44, 397,
	10, 44, 12, 44, 14, 44, 400, 11, 44, 3, 44, 5, 44, 403, 10, 44, 3, 45,
	3, 45, 7, 45, 407, 10, 45, 12, 45, 14, 45, 410, 11, 45, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	51, 3, 51, 5, 51, 437, 10, 51, 3, 52, 3, 52, 5, 52, 441, 10, 52, 3, 52,
	3, 52, 3, 52, 5, 52, 446, 10, 52, 3, 53, 3, 53, 3, 53, 3, 53, 5, 53, 452,
	10, 53, 3, 53, 3, 53, 3, 54, 5, 54, 457, 10, 54, 3, 54, 3, 54, 3, 54, 3,
	54, 3, 54, 5, 54, 464, 10, 54, 3, 55, 3, 55, 5, 55, 468, 10, 55, 3, 55,
	3, 55, 3, 56, 6, 56, 473, 10, 56, 13, 56, 14, 56, 474, 3, 57, 5, 57, 478,
	10, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 5, 57, 485, 10, 57, 3, 58, 6,
	58, 488, 10, 58, 13, 58, 14, 58, 489, 3, 59, 3, 59, 5, 59, 494, 10, 59,
	3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 5, 60, 503, 10, 60, 3,
	60, 5, 60, 506, 10, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 5, 60, 513,
	10, 60, 3, 61, 6, 61, 516, 10, 61, 13, 61, 14, 61, 517, 3, 61, 3, 61, 3,
	62, 3, 62, 5, 62, 524, 10, 62, 3, 62, 5, 62, 527, 10, 62, 3, 62, 3, 62,
	2, 2, 63, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29,
	57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38,
	75, 39, 77, 40, 79, 2, 81, 2, 83, 2, 85, 2, 87, 2, 89, 2, 91, 2, 93, 2,
	95, 2, 97, 2, 99, 2, 101, 2, 103, 2, 105, 2, 107, 2, 109, 2, 111, 2, 113,
	2, 115, 2, 117, 2, 119, 2, 121, 41, 123, 42, 3, 2, 19, 3, 2, 41, 41, 4,
	2, 36, 36, 94, 94, 5, 2, 36, 36, 41, 41, 94, 94, 5, 2, 78, 78, 87, 87,
	119, 119, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 68, 68, 100,
	100, 3, 2, 50, 51, 4, 2, 90, 90, 122, 122, 3, 2, 51, 59, 3, 2, 50, 57,
	5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47,
	47, 4, 2, 82, 82, 114, 114, 12, 2, 36, 36, 41, 41, 65, 65, 94, 94, 99,
	100, 104, 104, 112, 112, 116, 116, 118, 118, 120, 120, 4, 2, 11, 11, 34,
	34, 2, 564, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3,
	2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55,
	3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2,
	63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2,
	2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2,
	2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 3, 125, 3, 2, 2, 2, 5, 127,
	3, 2, 2, 2, 7, 129, 3, 2, 2, 2, 9, 131, 3, 2, 2, 2, 11, 133, 3, 2, 2, 2,
	13, 135, 3, 2, 2, 2, 15, 137, 3, 2, 2, 2, 17, 140, 3, 2, 2, 2, 19, 142,
	3, 2, 2, 2, 21, 145, 3, 2, 2, 2, 23, 148, 3, 2, 2, 2, 25, 159, 3, 2, 2,
	2, 27, 173, 3, 2, 2, 2, 29, 175, 3, 2, 2, 2, 31, 177, 3, 2, 2, 2, 33, 179,
	3, 2, 2, 2, 35, 181, 3, 2, 2, 2, 37, 183, 3, 2, 2, 2, 39, 185, 3, 2, 2,
	2, 41, 188, 3, 2, 2, 2, 43, 191, 3, 2, 2, 2, 45, 194, 3, 2, 2, 2, 47, 196,
	3, 2, 2, 2, 49, 198, 3, 2, 2, 2, 51, 205, 3, 2, 2, 2, 53, 211, 3, 2, 2,
	2, 55, 213, 3, 2, 2, 2, 57, 219, 3, 2, 2, 2, 59, 221, 3, 2, 2, 2, 61, 224,
	3, 2, 2, 2, 63, 231, 3, 2, 2, 2, 65, 267, 3, 2, 2, 2, 67, 296, 3, 2, 2,
	2, 69, 302, 3, 2, 2, 2, 71, 306, 3, 2, 2, 2, 73, 321, 3, 2, 2, 2, 75, 364,
	3, 2, 2, 2, 77, 366, 3, 2, 2, 2, 79, 381, 3, 2, 2, 2, 81, 383, 3, 2, 2,
	2, 83, 385, 3, 2, 2, 2, 85, 387, 3, 2, 2, 2, 87, 402, 3, 2, 2, 2, 89, 404,
	3, 2, 2, 2, 91, 411, 3, 2, 2, 2, 93, 415, 3, 2, 2, 2, 95, 417, 3, 2, 2,
	2, 97, 419, 3, 2, 2, 2, 99, 421, 3, 2, 2, 2, 101, 436, 3, 2, 2, 2, 103,
	445, 3, 2, 2, 2, 105, 447, 3, 2, 2, 2, 107, 463, 3, 2, 2, 2, 109, 465,
	3, 2, 2, 2, 111, 472, 3, 2, 2, 2, 113, 484, 3, 2, 2, 2, 115, 487, 3, 2,
	2, 2, 117, 491, 3, 2, 2, 2, 119, 512, 3, 2, 2, 2, 121, 515, 3, 2, 2, 2,
	123, 526, 3, 2, 2, 2, 125, 126, 7, 42, 2, 2, 126, 4, 3, 2, 2, 2, 127, 128,
	7, 43, 2, 2, 128, 6, 3, 2, 2, 2, 129, 130, 7, 93, 2, 2, 130, 8, 3, 2, 2,
	2, 131, 132, 7, 46, 2, 2, 132, 10, 3, 2, 2, 2, 133, 134, 7, 95, 2, 2, 134,
	12, 3, 2, 2, 2, 135, 136, 7, 62, 2, 2, 136, 14, 3, 2, 2, 2, 137, 138, 7,
	62, 2, 2, 138, 139, 7, 63, 2, 2, 139, 16, 3, 2, 2, 2, 140, 141, 7, 64,
	2, 2, 141, 18, 3, 2, 2, 2, 142, 143, 7, 64, 2, 2, 143, 144, 7, 63, 2, 2,
	144, 20, 3, 2, 2, 2, 145, 146, 7, 63, 2, 2, 146, 147, 7, 63, 2, 2, 147,
	22, 3, 2, 2, 2, 148, 149, 7, 35, 2, 2, 149, 150, 7, 63, 2, 2, 150, 24,
	3, 2, 2, 2, 151, 152, 7, 110, 2, 2, 152, 153, 7, 107, 2, 2, 153, 154, 7,
	109, 2, 2, 154, 160, 7, 103, 2, 2, 155, 156, 7, 78, 2, 2, 156, 157, 7,
	75, 2, 2, 157, 158, 7, 77, 2, 2, 158, 160, 7, 71, 2, 2, 159, 151, 3, 2,
	2, 2, 159, 155, 3, 2, 2, 2, 160, 26, 3, 2, 2, 2, 161, 162, 7, 103, 2, 2,
	162, 163, 7, 122, 2, 2, 163, 164, 7, 107, 2, 2, 164, 165, 7, 117, 2, 2,
	165, 166, 7, 118, 2, 2, 166, 174, 7, 117, 2, 2, 167, 168, 7, 71, 2, 2,
	168, 169, 7, 90, 2, 2, 169, 170, 7, 75, 2, 2, 170, 171, 7, 85, 2, 2, 171,
	172, 7, 86, 2, 2, 172, 174, 7, 85, 2, 2, 173, 161, 3, 2, 2, 2, 173, 167,
	3, 2, 2, 2, 174, 28, 3, 2, 2, 2, 175, 176, 7, 45, 2, 2, 176, 30, 3, 2,
	2, 2, 177, 178, 7, 47, 2, 2, 178, 32, 3, 2, 2, 2, 179, 180, 7, 44, 2, 2,
	180, 34, 3, 2, 2, 2, 181, 182, 7, 49, 2, 2, 182, 36, 3, 2, 2, 2, 183, 184,
	7, 39, 2, 2, 184, 38, 3, 2, 2, 2, 185, 186, 7, 44, 2, 2, 186, 187, 7, 44,
	2, 2, 187, 40, 3, 2, 2, 2, 188, 189, 7, 62, 2, 2, 189, 190, 7, 62, 2, 2,
	190, 42, 3, 2, 2, 2, 191, 192, 7, 64, 2, 2, 192, 193, 7, 64, 2, 2, 193,
	44, 3, 2, 2, 2, 194, 195, 7, 40, 2, 2, 195, 46, 3, 2, 2, 2, 196, 197, 7,
	126, 2, 2, 197, 48, 3, 2, 2, 2, 198, 199, 7, 96, 2, 2, 199, 50, 3, 2, 2,
	2, 200, 201, 7, 40, 2, 2, 201, 206, 7, 40, 2, 2, 202, 203, 7, 99, 2, 2,
	203, 204, 7, 112, 2, 2, 204, 206, 7, 102, 2, 2, 205, 200, 3, 2, 2, 2, 205,
	202, 3, 2, 2, 2, 206, 52, 3, 2, 2, 2, 207, 208, 7, 126, 2, 2, 208, 212,
	7, 126, 2, 2, 209, 210, 7, 113, 2, 2, 210, 212, 7, 116, 2, 2, 211, 207,
	3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 212, 54, 3, 2, 2, 2, 213, 214, 7, 128,
	2, 2, 214, 56, 3, 2, 2, 2, 215, 220, 7, 35, 2, 2, 216, 217, 7, 112, 2,
	2, 217, 218, 7, 113, 2, 2, 218, 220, 7, 118, 2, 2, 219, 215, 3, 2, 2, 2,
	219, 216, 3, 2, 2, 2, 220, 58, 3, 2, 2, 2, 221, 222, 7, 107, 2, 2, 222,
	223, 7, 112, 2, 2, 223, 60, 3, 2, 2, 2, 224, 225, 7, 112, 2, 2, 225, 226,
	7, 113, 2, 2, 226, 227, 7, 118, 2, 2, 227, 228, 7, 34, 2, 2, 228, 229,
	7, 107, 2, 2, 229, 230, 7, 112, 2, 2, 230, 62, 3, 2, 2, 2, 231, 236, 7,
	93, 2, 2, 232, 235, 5, 121, 61, 2, 233, 235, 5, 123, 62, 2, 234, 232, 3,
	2, 2, 2, 234, 233, 3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236, 234, 3, 2, 2,
	2, 236, 237, 3, 2, 2, 2, 237, 239, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 239,
	240, 7, 95, 2, 2, 240, 64, 3, 2, 2, 2, 241, 242, 7, 108, 2, 2, 242, 243,
	7, 117, 2, 2, 243, 244, 7, 113, 2, 2, 244, 245, 7, 112, 2, 2, 245, 246,
	7, 97, 2, 2, 246, 247, 7, 101, 2, 2, 247, 248, 7, 113, 2, 2, 248, 249,
	7, 112, 2, 2, 249, 250, 7, 118, 2, 2, 250, 251, 7, 99, 2, 2, 251, 252,
	7, 107, 2, 2, 252, 253, 7, 112, 2, 2, 253, 268, 7, 117, 2, 2, 254, 255,
	7, 76, 2, 2, 255, 256, 7, 85, 2, 2, 256, 257, 7, 81, 2, 2, 257, 258, 7,
	80, 2, 2, 258, 259, 7, 97, 2, 2, 259, 260, 7, 69, 2, 2, 260, 261, 7, 81,
	2, 2, 261, 262, 7, 80, 2, 2, 262, 263, 7, 86, 2, 2, 263, 264, 7, 67, 2,
	2, 264, 265, 7, 75, 2, 2, 265, 266, 7, 80, 2, 2, 266, 268, 7, 85, 2, 2,
	267, 241, 3, 2, 2, 2, 267, 254, 3, 2, 2, 2, 268, 66, 3, 2, 2, 2, 269, 270,
	7, 118, 2, 2, 270, 271, 7, 116, 2, 2, 271, 272, 7, 119, 2, 2, 272, 297,
	7, 103, 2, 2, 273, 274, 7, 86, 2, 2, 274, 275, 7, 116, 2, 2, 275, 276,
	7, 119, 2, 2, 276, 297, 7, 103, 2, 2, 277, 278, 7, 86, 2, 2, 278, 279,
	7, 84, 2, 2, 279, 280, 7, 87, 2, 2, 280, 297, 7, 71, 2, 2, 281, 282, 7,
	104, 2, 2, 282, 283, 7, 99, 2, 2, 283, 284, 7, 110, 2, 2, 284, 285, 7,
	117, 2, 2, 285, 297, 7, 103, 2, 2, 286, 287, 7, 72, 2, 2, 287, 288, 7,
	99, 2, 2, 288, 289, 7, 110, 2, 2, 289, 290, 7, 117, 2, 2, 290, 297, 7,
	103, 2, 2, 291, 292, 7, 72, 2, 2, 292, 293, 7, 67, 2, 2, 293, 294, 7, 78,
	2, 2, 294, 295, 7, 85, 2, 2, 295, 297, 7, 71, 2, 2, 296, 269, 3, 2, 2,
	2, 296, 273, 3, 2, 2, 2, 296, 277, 3, 2, 2, 2, 296, 281, 3, 2, 2, 2, 296,
	286, 3, 2, 2, 2, 296, 291, 3, 2, 2, 2, 297, 68, 3, 2, 2, 2, 298, 303, 5,
	87, 44, 2, 299, 303, 5, 89, 45, 2, 300, 303, 5, 91, 46, 2, 301, 303, 5,
	85, 43, 2, 302, 298, 3, 2, 2, 2, 302, 299, 3, 2, 2, 2, 302, 300, 3, 2,
	2, 2, 302, 301, 3, 2, 2, 2, 303, 70, 3, 2, 2, 2, 304, 307, 5, 103, 52,
	2, 305, 307, 5, 105, 53, 2, 306, 304, 3, 2, 2, 2, 306, 305, 3, 2, 2, 2,
	307, 72, 3, 2, 2, 2, 308, 313, 5, 81, 41, 2, 309, 312, 5, 81, 41, 2, 310,
	312, 5, 83, 42, 2, 311, 309, 3, 2, 2, 2, 311, 310, 3, 2, 2, 2, 312, 315,
	3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 322, 3, 2,
	2, 2, 315, 313, 3, 2, 2, 2, 316, 317, 7, 38, 2, 2, 317, 318, 7, 111, 2,
	2, 318, 319, 7, 103, 2, 2, 319, 320, 7, 118, 2, 2, 320, 322, 7, 99, 2,
	2, 321, 308, 3, 2, 2, 2, 321, 316, 3, 2, 2, 2, 322, 74, 3, 2, 2, 2, 323,
	325, 5, 79, 40, 2, 324, 323, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 325, 326,
	3, 2, 2, 2, 326, 327, 7, 36, 2, 2, 327, 339, 8, 38, 2, 2, 328, 333, 7,
	94, 2, 2, 329, 330, 7, 41, 2, 2, 330, 334, 8, 38, 3, 2, 331, 332, 10, 2,
	2, 2, 332, 334, 8, 38, 4, 2, 333, 329, 3, 2, 2, 2, 333, 331, 3, 2, 2, 2,
	334, 338, 3, 2, 2, 2, 335, 336, 10, 3, 2, 2, 336, 338, 8, 38, 5, 2, 337,
	328, 3, 2, 2, 2, 337, 335, 3, 2, 2, 2, 338, 341, 3, 2, 2, 2, 339, 337,
	3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 342, 3, 2, 2, 2, 341, 339, 3, 2,
	2, 2, 342, 343, 7, 36, 2, 2, 343, 365, 8, 38, 6, 2, 344, 345, 7, 41, 2,
	2, 345, 359, 8, 38, 7, 2, 346, 351, 7, 94, 2, 2, 347, 348, 7, 41, 2, 2,
	348, 352, 8, 38, 8, 2, 349, 350, 10, 2, 2, 2, 350, 352, 8, 38, 9, 2, 351,
	347, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 352, 358, 3, 2, 2, 2, 353, 354,
	7, 36, 2, 2, 354, 358, 8, 38, 10, 2, 355, 356, 10, 4, 2, 2, 356, 358, 8,
	38, 11, 2, 357, 346, 3, 2, 2, 2, 357, 353, 3, 2, 2, 2, 357, 355, 3, 2,
	2, 2, 358, 361, 3, 2, 2, 2, 359, 357, 3, 2, 2, 2, 359, 360, 3, 2, 2, 2,
	360, 362, 3, 2, 2, 2, 361, 359, 3, 2, 2, 2, 362, 363, 7, 41, 2, 2, 363,
	365, 8, 38, 12, 2, 364, 324, 3, 2, 2, 2, 364, 344, 3, 2, 2, 2, 365, 76,
	3, 2, 2, 2, 366, 374, 5, 73, 37, 2, 367, 370, 7, 93, 2, 2, 368, 371, 5,
	75, 38, 2, 369, 371, 5, 87, 44, 2, 370, 368, 3, 2, 2, 2, 370, 369, 3, 2,
	2, 2, 371, 372, 3, 2, 2, 2, 372, 373, 7, 95, 2, 2, 373, 375, 3, 2, 2, 2,
	374, 367, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 374, 3, 2, 2, 2, 376,
	377, 3, 2, 2, 2, 377, 78, 3, 2, 2, 2, 378, 379, 7, 119, 2, 2, 379, 382,
	7, 58, 2, 2, 380, 382, 9, 5, 2, 2, 381, 378, 3, 2, 2, 2, 381, 380, 3, 2,
	2, 2, 382, 80, 3, 2, 2, 2, 383, 384, 9, 6, 2, 2, 384, 82, 3, 2, 2, 2, 385,
	386, 9, 7, 2, 2, 386, 84, 3, 2, 2, 2, 387, 388, 7, 50, 2, 2, 388, 390,
	9, 8, 2, 2, 389, 391, 9, 9, 2, 2, 390, 389, 3, 2, 2, 2, 391, 392, 3, 2,
	2, 2, 392, 390, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2, 393, 86, 3, 2, 2, 2,
	394, 398, 5, 93, 47, 2, 395, 397, 5, 83, 42, 2, 396, 395, 3, 2, 2, 2, 397,
	400, 3, 2, 2, 2, 398, 396, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 403,
	3, 2, 2, 2, 400, 398, 3, 2, 2, 2, 401, 403, 7, 50, 2, 2, 402, 394, 3, 2,
	2, 2, 402, 401, 3, 2, 2, 2, 403, 88, 3, 2, 2, 2, 404, 408, 7, 50, 2, 2,
	405, 407, 5, 95, 48, 2, 406, 405, 3, 2, 2, 2, 407, 410, 3, 2, 2, 2, 408,
	406, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 90, 3, 2, 2, 2, 410, 408, 3,
	2, 2, 2, 411, 412, 7, 50, 2, 2, 412, 413, 9, 10, 2, 2, 413, 414, 5, 115,
	58, 2, 414, 92, 3, 2, 2, 2, 415, 416, 9, 11, 2, 2, 416, 94, 3, 2, 2, 2,
	417, 418, 9, 12, 2, 2, 418, 96, 3, 2, 2, 2, 419, 420, 9, 13, 2, 2, 420,
	98, 3, 2, 2, 2, 421, 422, 5, 97, 49, 2, 422, 423, 5, 97, 49, 2, 423, 424,
	5, 97, 49, 2, 424, 425, 5, 97, 49, 2, 425, 100, 3, 2, 2, 2, 426, 427, 7,
	94, 2, 2, 427, 428, 7, 119, 2, 2, 428, 429, 3, 2, 2, 2, 429, 437, 5, 99,
	50, 2, 430, 431, 7, 94, 2, 2, 431, 432, 7, 87, 2, 2, 432, 433, 3, 2, 2,
	2, 433, 434, 5, 99, 50, 2, 434, 435, 5, 99, 50, 2, 435, 437, 3, 2, 2, 2,
	436, 426, 3, 2, 2, 2, 436, 430, 3, 2, 2, 2, 437, 102, 3, 2, 2, 2, 438,
	440, 5, 107, 54, 2, 439, 441, 5, 109, 55, 2, 440, 439, 3, 2, 2, 2, 440,
	441, 3, 2, 2, 2, 441, 446, 3, 2, 2, 2, 442, 443, 5, 111, 56, 2, 443, 444,
	5, 109, 55, 2, 444, 446, 3, 2, 2, 2, 445, 438, 3, 2, 2, 2, 445, 442, 3,
	2, 2, 2, 446, 104, 3, 2, 2, 2, 447, 448, 7, 50, 2, 2, 448, 451, 9, 10,
	2, 2, 449, 452, 5, 113, 57, 2, 450, 452, 5, 115, 58, 2, 451, 449, 3, 2,
	2, 2, 451, 450, 3, 2, 2, 2, 452, 453, 3, 2, 2, 2, 453, 454, 5, 117, 59,
	2, 454, 106, 3, 2, 2, 2, 455, 457, 5, 111, 56, 2, 456, 455, 3, 2, 2, 2,
	456, 457, 3, 2, 2, 2, 457, 458, 3, 2, 2, 2, 458, 459, 7, 48, 2, 2, 459,
	464, 5, 111, 56, 2, 460, 461, 5, 111, 56, 2, 461, 462, 7, 48, 2, 2, 462,
	464, 3, 2, 2, 2, 463, 456, 3, 2, 2, 2, 463, 460, 3, 2, 2, 2, 464, 108,
	3, 2, 2, 2, 465, 467, 9, 14, 2, 2, 466, 468, 9, 15, 2, 2, 467, 466, 3,
	2, 2, 2, 467, 468, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 470, 5, 111,
	56, 2, 470, 110, 3, 2, 2, 2, 471, 473, 5, 83, 42, 2, 472, 471, 3, 2, 2,
	2, 473, 474, 3, 2, 2, 2, 474, 472, 3, 2, 2, 2, 474, 475, 3, 2, 2, 2, 475,
	112, 3, 2, 2, 2, 476, 478, 5, 115, 58, 2, 477, 476, 3, 2, 2, 2, 477, 478,
	3, 2, 2, 2, 478, 479, 3, 2, 2, 2, 479, 480, 7, 48, 2, 2, 480, 485, 5, 115,
	58, 2, 481, 482, 5, 115, 58, 2, 482, 483, 7, 48, 2, 2, 483, 485, 3, 2,
	2, 2, 484, 477, 3, 2, 2, 2, 484, 481, 3, 2, 2, 2, 485, 114, 3, 2, 2, 2,
	486, 488, 5, 97, 49, 2, 487, 486, 3, 2, 2, 2, 488, 489, 3, 2, 2, 2, 489,
	487, 3, 2, 2, 2, 489, 490, 3, 2, 2, 2, 490, 116, 3, 2, 2, 2, 491, 493,
	9, 16, 2, 2, 492, 494, 9, 15, 2, 2, 493, 492, 3, 2, 2, 2, 493, 494, 3,
	2, 2, 2, 494, 495, 3, 2, 2, 2, 495, 496, 5, 111, 56, 2, 496, 118, 3, 2,
	2, 2, 497, 498, 7, 94, 2, 2, 498, 513, 9, 17, 2, 2, 499, 500, 7, 94, 2,
	2, 500, 502, 5, 95, 48, 2, 501, 503, 5, 95, 48, 2, 502, 501, 3, 2, 2, 2,
	502, 503, 3, 2, 2, 2, 503, 505, 3, 2, 2, 2, 504, 506, 5, 95, 48, 2, 505,
	504, 3, 2, 2, 2, 505, 506, 3, 2, 2, 2, 506, 513, 3, 2, 2, 2, 507, 508,
	7, 94, 2, 2, 508, 509, 7, 122, 2, 2, 509, 510, 3, 2, 2, 2, 510, 513, 5,
	115, 58, 2, 511, 513, 5, 101, 51, 2, 512, 497, 3, 2, 2, 2, 512, 499, 3,
	2, 2, 2, 512, 507, 3, 2, 2, 2, 512, 511, 3, 2, 2, 2, 513, 120, 3, 2, 2,
	2, 514, 516, 9, 18, 2, 2, 515, 514, 3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517,
	515, 3, 2, 2, 2, 517, 518, 3, 2, 2, 2, 518, 519, 3, 2, 2, 2, 519, 520,
	8, 61, 13, 2, 520, 122, 3, 2, 2, 2, 521, 523, 7, 15, 2, 2, 522, 524, 7,
	12, 2, 2, 523, 522, 3, 2, 2, 2, 523, 524, 3, 2, 2, 2, 524, 527, 3, 2, 2,
	2, 525, 527, 7, 12, 2, 2, 526, 521, 3, 2, 2, 2, 526, 525, 3, 2, 2, 2, 527,
	528, 3, 2, 2, 2, 528, 529, 8, 62, 13, 2, 529, 124, 3, 2, 2, 2, 50, 2, 159,
	173, 205, 211, 219, 234, 236, 267, 296, 302, 306, 311, 313, 321, 324, 333,
	337, 339, 351, 357, 359, 364, 370, 376, 381, 392, 398, 402, 408, 436, 440,
	445, 451, 456, 463, 467, 474, 477, 484, 489, 493, 502, 505, 512, 517, 523,
	526, 14, 3, 38, 2, 3, 38, 3, 3, 38, 4, 3, 38, 5, 3, 38, 6, 3, 38, 7, 3,
	38, 8, 3, 38, 9, 3, 38, 10, 3, 38, 11, 3, 38, 12, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'['", "','", "']'", "'<'", "'<='", "'>'", "'>='", "'=='",
	"'!='", "", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'**'", "'<<'", "'>>'",
	"'&'", "'|'", "'^'", "", "", "'~'", "", "'in'", "'not in'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "LT", "LE", "GT", "GE", "EQ", "NE", "LIKE", "EXISTS",
	"ADD", "SUB", "MUL", "DIV", "MOD", "POW", "SHL", "SHR", "BAND", "BOR",
	"BXOR", "AND", "OR", "BNOT", "NOT", "IN", "NIN", "EmptyTerm", "JSONContains",
	"BooleanConstant", "IntegerConstant", "FloatingConstant", "Identifier",
	"StringLiteral", "JSONIdentifier", "Whitespace", "Newline",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "LT", "LE", "GT", "GE", "EQ", "NE",
	"LIKE", "EXISTS", "ADD", "SUB", "MUL", "DIV", "MOD", "POW", "SHL", "SHR",
	"BAND", "BOR", "BXOR", "AND", "OR", "BNOT", "NOT", "IN", "NIN", "EmptyTerm",
	"JSONContains", "BooleanConstant", "IntegerConstant", "FloatingConstant",
	"Identifier", "StringLiteral", "JSONIdentifier", "EncodingPrefix", "Nondigit",
	"Digit", "BinaryConstant", "DecimalConstant", "OctalConstant", "HexadecimalConstant",
	"NonzeroDigit", "OctalDigit", "HexadecimalDigit", "HexQuad", "UniversalCharacterName",
	"DecimalFloatingConstant", "HexadecimalFloatingConstant", "FractionalConstant",
	"ExponentPart", "DigitSequence", "HexadecimalFractionalConstant", "HexadecimalDigitSequence",
	"BinaryExponentPart", "EscapeSequence", "Whitespace", "Newline",
}

type PlanLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewPlanLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *PlanLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPlanLexer(input antlr.CharStream) *PlanLexer {
	l := new(PlanLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Plan.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PlanLexer tokens.
const (
	PlanLexerT__0             = 1
	PlanLexerT__1             = 2
	PlanLexerT__2             = 3
	PlanLexerT__3             = 4
	PlanLexerT__4             = 5
	PlanLexerLT               = 6
	PlanLexerLE               = 7
	PlanLexerGT               = 8
	PlanLexerGE               = 9
	PlanLexerEQ               = 10
	PlanLexerNE               = 11
	PlanLexerLIKE             = 12
	PlanLexerEXISTS           = 13
	PlanLexerADD              = 14
	PlanLexerSUB              = 15
	PlanLexerMUL              = 16
	PlanLexerDIV              = 17
	PlanLexerMOD              = 18
	PlanLexerPOW              = 19
	PlanLexerSHL              = 20
	PlanLexerSHR              = 21
	PlanLexerBAND             = 22
	PlanLexerBOR              = 23
	PlanLexerBXOR             = 24
	PlanLexerAND              = 25
	PlanLexerOR               = 26
	PlanLexerBNOT             = 27
	PlanLexerNOT              = 28
	PlanLexerIN               = 29
	PlanLexerNIN              = 30
	PlanLexerEmptyTerm        = 31
	PlanLexerJSONContains     = 32
	PlanLexerBooleanConstant  = 33
	PlanLexerIntegerConstant  = 34
	PlanLexerFloatingConstant = 35
	PlanLexerIdentifier       = 36
	PlanLexerStringLiteral    = 37
	PlanLexerJSONIdentifier   = 38
	PlanLexerWhitespace       = 39
	PlanLexerNewline          = 40
)

var str = ""

func (l *PlanLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 36:
		l.StringLiteral_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *PlanLexer) StringLiteral_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
		str += "\""

	case 1:
		str += "'"

	case 2:
		str += ("\\" + string(l.GetInputStream().LA(-1)))

	case 3:
		str += string(l.GetInputStream().LA(-1))

	case 4:

		str += "\""
		l.SetText(str)
		str = ""

	case 5:
		str += "'"

	case 6:
		str += "'"

	case 7:
		str += ("\\" + string(l.GetInputStream().LA(-1)))

	case 8:
		str += "\\\""

	case 9:
		str += string(l.GetInputStream().LA(-1))

	case 10:

		str += "'"
		l.SetText(str)
		str = ""

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
