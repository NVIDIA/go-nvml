// Copyright (c) 2019, NVIDIA CORPORATION. All rights reserved.

package nvml

// nvml.EventSetCreate()
func EventSetCreate() (EventSet, Return) {
	var Set EventSet
	ret := nvmlEventSetCreate(&Set)
	return Set, ret
}

// nvml.EventSetWait()
func EventSetWait(Set EventSet, Timeoutms uint32) (EventData, Return) {
	var Data EventData
	ret := nvmlEventSetWait(Set, &Data, Timeoutms)
	return Data, ret
}

func (Set EventSet) Wait(Timeoutms uint32) (EventData, Return) {
	return EventSetWait(Set, Timeoutms)
}

// nvml.EventSetFree()
func EventSetFree(Set EventSet) Return {
	return nvmlEventSetFree(Set)
}

func (Set EventSet) Free() Return {
	return EventSetFree(Set)
}
