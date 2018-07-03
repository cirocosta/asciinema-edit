package editor

// Quantize constraints a set of inputs that lie in a range to a single
// value that corresponds to the lower bound of such range.
//
// 1. capture all delays
// 2. for each delay, check if the delay fits in the quantization
//    range.
// 3. if it fits, reduce the delay to the maximum allowed (floor of
//    the quantization range).
// 4. adjust the rest of the event stream.
