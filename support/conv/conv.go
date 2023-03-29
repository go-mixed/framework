package conv

func ConvertArgs1[A1 any](args ...any) (a1 A1) {
	a1, _, _, _, _, _, _, _, _, _ = ConvertArgs10[A1, any, any, any, any, any, any, any, any, any](args...)
	return
}

func ConvertArgs2[A1 any, A2 any](args ...any) (a1 A1, a2 A2) {
	a1, a2, _, _, _, _, _, _, _, _ = ConvertArgs10[A1, A2, any, any, any, any, any, any, any, any](args...)
	return
}

func ConvertArgs3[A1 any, A2 any, A3 any](args ...any) (a1 A1, a2 A2, a3 A3) {
	a1, a2, a3, _, _, _, _, _, _, _ = ConvertArgs10[A1, A2, A3, any, any, any, any, any, any, any](args...)
	return
}
func ConvertArgs4[A1 any, A2 any, A3 any, A4 any](args ...any) (a1 A1, a2 A2, a3 A3, a4 A4) {
	a1, a2, a3, a4, _, _, _, _, _, _ = ConvertArgs10[A1, A2, A3, A4, any, any, any, any, any, any](args...)
	return
}

func ConvertArgs5[A1 any, A2 any, A3 any, A4 any, A5 any](args ...any) (a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) {
	a1, a2, a3, a4, a5, _, _, _, _, _ = ConvertArgs10[A1, A2, A3, A4, A5, any, any, any, any, any](args...)
	return
}

func ConvertArgs6[A1 any, A2 any, A3 any, A4 any, A5 any, A6 any](args ...any) (a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) {
	a1, a2, a3, a4, a5, a6, _, _, _, _ = ConvertArgs10[A1, A2, A3, A4, A5, A6, any, any, any, any](args...)
	return
}

func ConvertArgs7[A1 any, A2 any, A3 any, A4 any, A5 any, A6 any, A7 any](args ...any) (a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) {
	a1, a2, a3, a4, a5, a6, a7, _, _, _ = ConvertArgs10[A1, A2, A3, A4, A5, A6, A7, any, any, any](args...)
	return
}

func ConvertArgs8[A1 any, A2 any, A3 any, A4 any, A5 any, A6 any, A7 any, A8 any](args ...any) (a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) {
	a1, a2, a3, a4, a5, a6, a7, a8, _, _ = ConvertArgs10[A1, A2, A3, A4, A5, A6, A7, A8, any, any](args...)
	return
}

func ConvertArgs9[A1 any, A2 any, A3 any, A4 any, A5 any, A6 any, A7 any, A8 any, A9 any](args ...any) (a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) {
	a1, a2, a3, a4, a5, a6, a7, a8, a9, _ = ConvertArgs10[A1, A2, A3, A4, A5, A6, A7, A8, A9, any](args...)
	return
}

func ConvertArgs10[A1 any, A2 any, A3 any, A4 any, A5 any, A6 any, A7 any, A8 any, A9 any, A10 any](args ...any) (a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) {
	l := len(args)
	if l >= 1 {
		a1 = args[0].(A1)
	}
	if l >= 2 {
		a2 = args[1].(A2)
	}
	if l >= 3 {
		a3 = args[2].(A3)
	}
	if l >= 4 {
		a4 = args[3].(A4)
	}
	if l >= 5 {
		a5 = args[4].(A5)
	}
	if l >= 6 {
		a6 = args[5].(A6)
	}
	if l >= 7 {
		a7 = args[6].(A7)
	}
	if l >= 8 {
		a8 = args[7].(A8)
	}
	if l >= 9 {
		a9 = args[8].(A9)
	}
	if l >= 10 {
		a10 = args[9].(A10)
	}

	return
}
