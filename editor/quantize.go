package editor

// Quantize constraints a set of inputs that lie in a range to a single
// value that corresponds to the lower bound of such range.
//
// For instance, consider the following timestamps:
//
//	 1  2  5  9 10 11 12
//
// Assuming that we quantize over [2,6), we'd cut any delays between 2 and
// 6 seconds to 2 second:
//
//	 1  2  4  6  7  8  9	(with times already adjusted)
//
// The euristic is fairly simple:
//
// 1. capture all delays
// 2. for each delay, check if it's within an acceptable delay range
// 3. if it fits, reduce the delay to the maximum allowed (floor of
//    the quantization range).
// 4. adjust the rest of the event stream.
