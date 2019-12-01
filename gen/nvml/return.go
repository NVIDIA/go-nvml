// Copyright (c) 2019, NVIDIA CORPORATION. All rights reserved.

package nvml

// nvml.ErrorString()
func ErrorString(Result Return) string {
	return nvmlErrorString(Result)
}

func (r Return) String() string {
	return ErrorString(r)
}

func (r Return) Error() string {
	return ErrorString(r)
}
