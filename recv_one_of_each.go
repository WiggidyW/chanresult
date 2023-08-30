package chanresult

type Recv2Data[
	T1 any,
	T2 any,
] struct {
	T1 T1
	T2 T2
}

type Recv3Data[
	T1 any,
	T2 any,
	T3 any,
] struct {
	T1 T1
	T2 T2
	T3 T3
}

type Recv4Data[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
}

type Recv5Data[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
	T5 T5
}

type Recv6Data[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
	T5 T5
	T6 T6
}

type Recv7Data[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
	T5 T5
	T6 T6
	T7 T7
}

type Recv8Data[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
	T5 T5
	T6 T6
	T7 T7
	T8 T8
}

type Recv9Data[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
	T9 any,
] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
	T5 T5
	T6 T6
	T7 T7
	T8 T8
	T9 T9
}

type Recv10Data[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
	T9 any,
	T10 any,
] struct {
	T1  T1
	T2  T2
	T3  T3
	T4  T4
	T5  T5
	T6  T6
	T7  T7
	T8  T8
	T9  T9
	T10 T10
}

func RecvOneOfEach2[
	T1 any,
	T2 any,
](
	chn1 ChanRecvResult[T1],
	chn2 ChanRecvResult[T2],
) (
	data *Recv2Data[
		T1,
		T2,
	],
	err error,
) {
	select {
	case data.T1 = <-chn1.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	}

	select {
	case data.T2 = <-chn2.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	}

	return data, nil
}
func RecvOneOfEach3[
	T1 any,
	T2 any,
	T3 any,
](
	chn1 ChanRecvResult[T1],
	chn2 ChanRecvResult[T2],
	chn3 ChanRecvResult[T3],
) (
	data *Recv3Data[
		T1,
		T2,
		T3,
	],
	err error,
) {
	data = &Recv3Data[T1, T2, T3]{}

	select {
	case data.T1 = <-chn1.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	}

	select {
	case data.T2 = <-chn2.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	}

	select {
	case data.T3 = <-chn3.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	}

	return data, nil
}

func RecvOneOfEach4[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
](
	chn1 ChanRecvResult[T1],
	chn2 ChanRecvResult[T2],
	chn3 ChanRecvResult[T3],
	chn4 ChanRecvResult[T4],
) (
	data *Recv4Data[
		T1,
		T2,
		T3,
		T4,
	],
	err error,
) {
	data = &Recv4Data[T1, T2, T3, T4]{}

	select {
	case data.T1 = <-chn1.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	}

	select {
	case data.T2 = <-chn2.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	}

	select {
	case data.T3 = <-chn3.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	}

	select {
	case data.T4 = <-chn4.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	}

	return data, nil
}

func RecvOneOfEach5[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
](
	chn1 ChanRecvResult[T1],
	chn2 ChanRecvResult[T2],
	chn3 ChanRecvResult[T3],
	chn4 ChanRecvResult[T4],
	chn5 ChanRecvResult[T5],
) (
	data *Recv5Data[
		T1,
		T2,
		T3,
		T4,
		T5,
	],
	err error,
) {
	data = &Recv5Data[T1, T2, T3, T4, T5]{}

	select {
	case data.T1 = <-chn1.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	}

	select {
	case data.T2 = <-chn2.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	}

	select {
	case data.T3 = <-chn3.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	}

	select {
	case data.T4 = <-chn4.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	}

	select {
	case data.T5 = <-chn5.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	}

	return data, nil
}

func RecvOneOfEach6[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
](
	chn1 ChanRecvResult[T1],
	chn2 ChanRecvResult[T2],
	chn3 ChanRecvResult[T3],
	chn4 ChanRecvResult[T4],
	chn5 ChanRecvResult[T5],
	chn6 ChanRecvResult[T6],
) (
	data *Recv6Data[
		T1,
		T2,
		T3,
		T4,
		T5,
		T6,
	],
	err error,
) {
	data = &Recv6Data[T1, T2, T3, T4, T5, T6]{}

	select {
	case data.T1 = <-chn1.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	}

	select {
	case data.T2 = <-chn2.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	}

	select {
	case data.T3 = <-chn3.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	}

	select {
	case data.T4 = <-chn4.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	}

	select {
	case data.T5 = <-chn5.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	}

	select {
	case data.T6 = <-chn6.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	}

	return data, nil
}
func RecvOneOfEach7[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
](
	chn1 ChanRecvResult[T1],
	chn2 ChanRecvResult[T2],
	chn3 ChanRecvResult[T3],
	chn4 ChanRecvResult[T4],
	chn5 ChanRecvResult[T5],
	chn6 ChanRecvResult[T6],
	chn7 ChanRecvResult[T7],
) (
	data *Recv7Data[
		T1,
		T2,
		T3,
		T4,
		T5,
		T6,
		T7,
	],
	err error,
) {
	data = &Recv7Data[T1, T2, T3, T4, T5, T6, T7]{}

	select {
	case data.T1 = <-chn1.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	}

	select {
	case data.T2 = <-chn2.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	}

	select {
	case data.T3 = <-chn3.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	}

	select {
	case data.T4 = <-chn4.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	}

	select {
	case data.T5 = <-chn5.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	}

	select {
	case data.T6 = <-chn6.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	}

	select {
	case data.T7 = <-chn7.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	}

	return data, nil
}
func RecvOneOfEach8[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
](
	chn1 ChanRecvResult[T1],
	chn2 ChanRecvResult[T2],
	chn3 ChanRecvResult[T3],
	chn4 ChanRecvResult[T4],
	chn5 ChanRecvResult[T5],
	chn6 ChanRecvResult[T6],
	chn7 ChanRecvResult[T7],
	chn8 ChanRecvResult[T8],
) (
	data *Recv8Data[
		T1,
		T2,
		T3,
		T4,
		T5,
		T6,
		T7,
		T8,
	],
	err error,
) {
	data = &Recv8Data[T1, T2, T3, T4, T5, T6, T7, T8]{}

	select {
	case data.T1 = <-chn1.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	}

	select {
	case data.T2 = <-chn2.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	}

	select {
	case data.T3 = <-chn3.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	}

	select {
	case data.T4 = <-chn4.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	}

	select {
	case data.T5 = <-chn5.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	}

	select {
	case data.T6 = <-chn6.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	}

	select {
	case data.T7 = <-chn7.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	}

	select {
	case data.T8 = <-chn8.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	}

	return data, nil
}
func RecvOneOfEach9[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
	T9 any,
](
	chn1 ChanRecvResult[T1],
	chn2 ChanRecvResult[T2],
	chn3 ChanRecvResult[T3],
	chn4 ChanRecvResult[T4],
	chn5 ChanRecvResult[T5],
	chn6 ChanRecvResult[T6],
	chn7 ChanRecvResult[T7],
	chn8 ChanRecvResult[T8],
	chn9 ChanRecvResult[T9],
) (
	data *Recv9Data[
		T1,
		T2,
		T3,
		T4,
		T5,
		T6,
		T7,
		T8,
		T9,
	],
	err error,
) {
	data = &Recv9Data[T1, T2, T3, T4, T5, T6, T7, T8, T9]{}

	select {
	case data.T1 = <-chn1.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	}

	select {
	case data.T2 = <-chn2.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	}

	select {
	case data.T3 = <-chn3.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	}

	select {
	case data.T4 = <-chn4.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	}

	select {
	case data.T5 = <-chn5.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	}

	select {
	case data.T6 = <-chn6.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	}

	select {
	case data.T7 = <-chn7.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	}

	select {
	case data.T8 = <-chn8.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	}

	select {
	case data.T9 = <-chn9.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	}

	return data, nil
}
func RecvOneOfEach10[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
	T9 any,
	T10 any,
](
	chn1 ChanRecvResult[T1],
	chn2 ChanRecvResult[T2],
	chn3 ChanRecvResult[T3],
	chn4 ChanRecvResult[T4],
	chn5 ChanRecvResult[T5],
	chn6 ChanRecvResult[T6],
	chn7 ChanRecvResult[T7],
	chn8 ChanRecvResult[T8],
	chn9 ChanRecvResult[T9],
	chn10 ChanRecvResult[T10],
) (
	data *Recv10Data[
		T1,
		T2,
		T3,
		T4,
		T5,
		T6,
		T7,
		T8,
		T9,
		T10,
	],
	err error,
) {
	data = &Recv10Data[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{}

	select {
	case data.T1 = <-chn1.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case err = <-chn10.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	case <-chn10.inner.ctx.Done():
		return nil, chn10.inner.ctx.Err()
	}

	select {
	case data.T2 = <-chn2.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case err = <-chn10.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	case <-chn10.inner.ctx.Done():
		return nil, chn10.inner.ctx.Err()
	}

	select {
	case data.T3 = <-chn3.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case err = <-chn10.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	case <-chn10.inner.ctx.Done():
		return nil, chn10.inner.ctx.Err()
	}

	select {
	case data.T4 = <-chn4.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case err = <-chn10.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	case <-chn10.inner.ctx.Done():
		return nil, chn10.inner.ctx.Err()
	}

	select {
	case data.T5 = <-chn5.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case err = <-chn10.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	case <-chn10.inner.ctx.Done():
		return nil, chn10.inner.ctx.Err()
	}

	select {
	case data.T6 = <-chn6.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case err = <-chn10.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	case <-chn10.inner.ctx.Done():
		return nil, chn10.inner.ctx.Err()
	}

	select {
	case data.T7 = <-chn7.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case err = <-chn10.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	case <-chn10.inner.ctx.Done():
		return nil, chn10.inner.ctx.Err()
	}

	select {
	case data.T8 = <-chn8.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case err = <-chn10.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	case <-chn10.inner.ctx.Done():
		return nil, chn10.inner.ctx.Err()
	}

	select {
	case data.T9 = <-chn9.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case err = <-chn10.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	case <-chn10.inner.ctx.Done():
		return nil, chn10.inner.ctx.Err()
	}

	select {
	case data.T10 = <-chn10.inner.chnOk:
	case err = <-chn1.inner.chnErr:
		return nil, err
	case err = <-chn2.inner.chnErr:
		return nil, err
	case err = <-chn3.inner.chnErr:
		return nil, err
	case err = <-chn4.inner.chnErr:
		return nil, err
	case err = <-chn5.inner.chnErr:
		return nil, err
	case err = <-chn6.inner.chnErr:
		return nil, err
	case err = <-chn7.inner.chnErr:
		return nil, err
	case err = <-chn8.inner.chnErr:
		return nil, err
	case err = <-chn9.inner.chnErr:
		return nil, err
	case err = <-chn10.inner.chnErr:
		return nil, err
	case <-chn1.inner.ctx.Done():
		return nil, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return nil, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return nil, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return nil, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return nil, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return nil, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return nil, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return nil, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return nil, chn9.inner.ctx.Err()
	case <-chn10.inner.ctx.Done():
		return nil, chn10.inner.ctx.Err()
	}

	return data, nil
}
