package main

func MonitorGenVar(varName string) int {
	switch varName {
	case "a":
		return GetValMem(a)
	case "b":
		return GetValMem(b)
	case "c":
		return GetValMem(c)
	}
	return GetValMem(a, b, c)
}
