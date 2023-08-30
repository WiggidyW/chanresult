package chanresult

type RecvVariant uint8

const (
	RECV_1 RecvVariant = iota
	RECV_2
	RECV_3
	RECV_4
	RECV_5
	RECV_6
	RECV_7
	RECV_8
	RECV_9
	RECV_10
)

func Recv2[
	T1 any,
	T2 any,
](
	chn1 ChanRecvResult[T1],
	chn2 ChanRecvResult[T2],
) (
	variant RecvVariant,
	t1 T1,
	t2 T2,
	err error,
) {
	select {
	case t1 = <-chn1.inner.chnOk:
		return RECV_1, t1, t2, nil
	case t2 = <-chn2.inner.chnOk:
		return RECV_2, t1, t2, nil

	case err = <-chn1.inner.chnErr:
		return RECV_1, t1, t2, err
	case err = <-chn2.inner.chnErr:
		return RECV_2, t1, t2, err

	case <-chn1.inner.ctx.Done():
		return RECV_1, t1, t2, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return RECV_2, t1, t2, chn2.inner.ctx.Err()
	}
}

func Recv3[
	T1 any,
	T2 any,
	T3 any,
](
	chn1 ChanRecvResult[T1],
	chn2 ChanRecvResult[T2],
	chn3 ChanRecvResult[T3],
) (
	variant RecvVariant,
	t1 T1,
	t2 T2,
	t3 T3,
	err error,
) {
	select {
	case t1 = <-chn1.inner.chnOk:
		return RECV_1, t1, t2, t3, nil
	case t2 = <-chn2.inner.chnOk:
		return RECV_2, t1, t2, t3, nil
	case t3 = <-chn3.inner.chnOk:
		return RECV_3, t1, t2, t3, nil

	case err = <-chn1.inner.chnErr:
		return RECV_1, t1, t2, t3, err
	case err = <-chn2.inner.chnErr:
		return RECV_2, t1, t2, t3, err
	case err = <-chn3.inner.chnErr:
		return RECV_3, t1, t2, t3, err

	case <-chn1.inner.ctx.Done():
		return RECV_1, t1, t2, t3, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return RECV_2, t1, t2, t3, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return RECV_3, t1, t2, t3, chn3.inner.ctx.Err()
	}
}

func Recv4[
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
	variant RecvVariant,
	t1 T1,
	t2 T2,
	t3 T3,
	t4 T4,
	err error,
) {
	select {
	case t1 = <-chn1.inner.chnOk:
		return RECV_1, t1, t2, t3, t4, nil
	case t2 = <-chn2.inner.chnOk:
		return RECV_2, t1, t2, t3, t4, nil
	case t3 = <-chn3.inner.chnOk:
		return RECV_3, t1, t2, t3, t4, nil
	case t4 = <-chn4.inner.chnOk:
		return RECV_4, t1, t2, t3, t4, nil

	case err = <-chn1.inner.chnErr:
		return RECV_1, t1, t2, t3, t4, err
	case err = <-chn2.inner.chnErr:
		return RECV_2, t1, t2, t3, t4, err
	case err = <-chn3.inner.chnErr:
		return RECV_3, t1, t2, t3, t4, err
	case err = <-chn4.inner.chnErr:
		return RECV_4, t1, t2, t3, t4, err

	case <-chn1.inner.ctx.Done():
		return RECV_1, t1, t2, t3, t4, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return RECV_2, t1, t2, t3, t4, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return RECV_3, t1, t2, t3, t4, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return RECV_4, t1, t2, t3, t4, chn4.inner.ctx.Err()
	}
}

func Recv5[
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
	variant RecvVariant,
	t1 T1,
	t2 T2,
	t3 T3,
	t4 T4,
	t5 T5,
	err error,
) {
	select {
	case t1 = <-chn1.inner.chnOk:
		return RECV_1, t1, t2, t3, t4, t5, nil
	case t2 = <-chn2.inner.chnOk:
		return RECV_2, t1, t2, t3, t4, t5, nil
	case t3 = <-chn3.inner.chnOk:
		return RECV_3, t1, t2, t3, t4, t5, nil
	case t4 = <-chn4.inner.chnOk:
		return RECV_4, t1, t2, t3, t4, t5, nil
	case t5 = <-chn5.inner.chnOk:
		return RECV_5, t1, t2, t3, t4, t5, nil

	case err = <-chn1.inner.chnErr:
		return RECV_1, t1, t2, t3, t4, t5, err
	case err = <-chn2.inner.chnErr:
		return RECV_2, t1, t2, t3, t4, t5, err
	case err = <-chn3.inner.chnErr:
		return RECV_3, t1, t2, t3, t4, t5, err
	case err = <-chn4.inner.chnErr:
		return RECV_4, t1, t2, t3, t4, t5, err
	case err = <-chn5.inner.chnErr:
		return RECV_5, t1, t2, t3, t4, t5, err

	case <-chn1.inner.ctx.Done():
		return RECV_1, t1, t2, t3, t4, t5, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return RECV_2, t1, t2, t3, t4, t5, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return RECV_3, t1, t2, t3, t4, t5, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return RECV_4, t1, t2, t3, t4, t5, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return RECV_5, t1, t2, t3, t4, t5, chn5.inner.ctx.Err()
	}
}

func Recv6[
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
	variant RecvVariant,
	t1 T1,
	t2 T2,
	t3 T3,
	t4 T4,
	t5 T5,
	t6 T6,
	err error,
) {
	select {
	case t1 = <-chn1.inner.chnOk:
		return RECV_1, t1, t2, t3, t4, t5, t6, nil
	case t2 = <-chn2.inner.chnOk:
		return RECV_2, t1, t2, t3, t4, t5, t6, nil
	case t3 = <-chn3.inner.chnOk:
		return RECV_3, t1, t2, t3, t4, t5, t6, nil
	case t4 = <-chn4.inner.chnOk:
		return RECV_4, t1, t2, t3, t4, t5, t6, nil
	case t5 = <-chn5.inner.chnOk:
		return RECV_5, t1, t2, t3, t4, t5, t6, nil
	case t6 = <-chn6.inner.chnOk:
		return RECV_6, t1, t2, t3, t4, t5, t6, nil

	case err = <-chn1.inner.chnErr:
		return RECV_1, t1, t2, t3, t4, t5, t6, err
	case err = <-chn2.inner.chnErr:
		return RECV_2, t1, t2, t3, t4, t5, t6, err
	case err = <-chn3.inner.chnErr:
		return RECV_3, t1, t2, t3, t4, t5, t6, err
	case err = <-chn4.inner.chnErr:
		return RECV_4, t1, t2, t3, t4, t5, t6, err
	case err = <-chn5.inner.chnErr:
		return RECV_5, t1, t2, t3, t4, t5, t6, err
	case err = <-chn6.inner.chnErr:
		return RECV_6, t1, t2, t3, t4, t5, t6, err

	case <-chn1.inner.ctx.Done():
		return RECV_1, t1, t2, t3, t4, t5, t6, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return RECV_2, t1, t2, t3, t4, t5, t6, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return RECV_3, t1, t2, t3, t4, t5, t6, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return RECV_4, t1, t2, t3, t4, t5, t6, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return RECV_5, t1, t2, t3, t4, t5, t6, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return RECV_6, t1, t2, t3, t4, t5, t6, chn6.inner.ctx.Err()
	}
}

func Recv7[
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
	variant RecvVariant,
	t1 T1,
	t2 T2,
	t3 T3,
	t4 T4,
	t5 T5,
	t6 T6,
	t7 T7,
	err error,
) {
	select {
	case t1 = <-chn1.inner.chnOk:
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, nil
	case t2 = <-chn2.inner.chnOk:
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, nil
	case t3 = <-chn3.inner.chnOk:
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, nil
	case t4 = <-chn4.inner.chnOk:
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, nil
	case t5 = <-chn5.inner.chnOk:
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, nil
	case t6 = <-chn6.inner.chnOk:
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, nil
	case t7 = <-chn7.inner.chnOk:
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, nil

	case err = <-chn1.inner.chnErr:
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, err
	case err = <-chn2.inner.chnErr:
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, err
	case err = <-chn3.inner.chnErr:
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, err
	case err = <-chn4.inner.chnErr:
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, err
	case err = <-chn5.inner.chnErr:
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, err
	case err = <-chn6.inner.chnErr:
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, err
	case err = <-chn7.inner.chnErr:
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, err

	case <-chn1.inner.ctx.Done():
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, chn7.inner.ctx.Err()
	}
}

func Recv8[
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
	variant RecvVariant,
	t1 T1,
	t2 T2,
	t3 T3,
	t4 T4,
	t5 T5,
	t6 T6,
	t7 T7,
	t8 T8,
	err error,
) {
	select {
	case t1 = <-chn1.inner.chnOk:
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, t8, nil
	case t2 = <-chn2.inner.chnOk:
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, t8, nil
	case t3 = <-chn3.inner.chnOk:
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, t8, nil
	case t4 = <-chn4.inner.chnOk:
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, t8, nil
	case t5 = <-chn5.inner.chnOk:
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, t8, nil
	case t6 = <-chn6.inner.chnOk:
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, t8, nil
	case t7 = <-chn7.inner.chnOk:
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, t8, nil
	case t8 = <-chn8.inner.chnOk:
		return RECV_8, t1, t2, t3, t4, t5, t6, t7, t8, nil

	case err = <-chn1.inner.chnErr:
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, t8, err
	case err = <-chn2.inner.chnErr:
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, t8, err
	case err = <-chn3.inner.chnErr:
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, t8, err
	case err = <-chn4.inner.chnErr:
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, t8, err
	case err = <-chn5.inner.chnErr:
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, t8, err
	case err = <-chn6.inner.chnErr:
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, t8, err
	case err = <-chn7.inner.chnErr:
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, t8, err
	case err = <-chn8.inner.chnErr:
		return RECV_8, t1, t2, t3, t4, t5, t6, t7, t8, err

	case <-chn1.inner.ctx.Done():
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, t8, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, t8, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, t8, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, t8, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, t8, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, t8, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, t8, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return RECV_8, t1, t2, t3, t4, t5, t6, t7, t8, chn8.inner.ctx.Err()
	}
}

func Recv9[
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
	variant RecvVariant,
	t1 T1,
	t2 T2,
	t3 T3,
	t4 T4,
	t5 T5,
	t6 T6,
	t7 T7,
	t8 T8,
	t9 T9,
	err error,
) {
	select {
	case t1 = <-chn1.inner.chnOk:
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, t8, t9, nil
	case t2 = <-chn2.inner.chnOk:
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, t8, t9, nil
	case t3 = <-chn3.inner.chnOk:
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, t8, t9, nil
	case t4 = <-chn4.inner.chnOk:
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, t8, t9, nil
	case t5 = <-chn5.inner.chnOk:
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, t8, t9, nil
	case t6 = <-chn6.inner.chnOk:
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, t8, t9, nil
	case t7 = <-chn7.inner.chnOk:
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, t8, t9, nil
	case t8 = <-chn8.inner.chnOk:
		return RECV_8, t1, t2, t3, t4, t5, t6, t7, t8, t9, nil
	case t9 = <-chn9.inner.chnOk:
		return RECV_9, t1, t2, t3, t4, t5, t6, t7, t8, t9, nil

	case err = <-chn1.inner.chnErr:
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, t8, t9, err
	case err = <-chn2.inner.chnErr:
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, t8, t9, err
	case err = <-chn3.inner.chnErr:
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, t8, t9, err
	case err = <-chn4.inner.chnErr:
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, t8, t9, err
	case err = <-chn5.inner.chnErr:
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, t8, t9, err
	case err = <-chn6.inner.chnErr:
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, t8, t9, err
	case err = <-chn7.inner.chnErr:
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, t8, t9, err
	case err = <-chn8.inner.chnErr:
		return RECV_8, t1, t2, t3, t4, t5, t6, t7, t8, t9, err
	case err = <-chn9.inner.chnErr:
		return RECV_9, t1, t2, t3, t4, t5, t6, t7, t8, t9, err

	case <-chn1.inner.ctx.Done():
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, t8, t9, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, t8, t9, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, t8, t9, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, t8, t9, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, t8, t9, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, t8, t9, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, t8, t9, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return RECV_8, t1, t2, t3, t4, t5, t6, t7, t8, t9, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return RECV_9, t1, t2, t3, t4, t5, t6, t7, t8, t9, chn9.inner.ctx.Err()
	}
}

func Recv10[
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
	variant RecvVariant,
	t1 T1,
	t2 T2,
	t3 T3,
	t4 T4,
	t5 T5,
	t6 T6,
	t7 T7,
	t8 T8,
	t9 T9,
	t10 T10,
	err error,
) {
	select {
	case t1 = <-chn1.inner.chnOk:
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, nil
	case t2 = <-chn2.inner.chnOk:
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, nil
	case t3 = <-chn3.inner.chnOk:
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, nil
	case t4 = <-chn4.inner.chnOk:
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, nil
	case t5 = <-chn5.inner.chnOk:
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, nil
	case t6 = <-chn6.inner.chnOk:
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, nil
	case t7 = <-chn7.inner.chnOk:
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, nil
	case t8 = <-chn8.inner.chnOk:
		return RECV_8, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, nil
	case t9 = <-chn9.inner.chnOk:
		return RECV_9, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, nil
	case t10 = <-chn10.inner.chnOk:
		return RECV_10, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, nil

	case err = <-chn1.inner.chnErr:
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, err
	case err = <-chn2.inner.chnErr:
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, err
	case err = <-chn3.inner.chnErr:
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, err
	case err = <-chn4.inner.chnErr:
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, err
	case err = <-chn5.inner.chnErr:
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, err
	case err = <-chn6.inner.chnErr:
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, err
	case err = <-chn7.inner.chnErr:
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, err
	case err = <-chn8.inner.chnErr:
		return RECV_8, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, err
	case err = <-chn9.inner.chnErr:
		return RECV_9, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, err
	case err = <-chn10.inner.chnErr:
		return RECV_10, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, err

	case <-chn1.inner.ctx.Done():
		return RECV_1, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, chn1.inner.ctx.Err()
	case <-chn2.inner.ctx.Done():
		return RECV_2, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, chn2.inner.ctx.Err()
	case <-chn3.inner.ctx.Done():
		return RECV_3, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, chn3.inner.ctx.Err()
	case <-chn4.inner.ctx.Done():
		return RECV_4, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, chn4.inner.ctx.Err()
	case <-chn5.inner.ctx.Done():
		return RECV_5, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, chn5.inner.ctx.Err()
	case <-chn6.inner.ctx.Done():
		return RECV_6, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, chn6.inner.ctx.Err()
	case <-chn7.inner.ctx.Done():
		return RECV_7, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, chn7.inner.ctx.Err()
	case <-chn8.inner.ctx.Done():
		return RECV_8, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, chn8.inner.ctx.Err()
	case <-chn9.inner.ctx.Done():
		return RECV_9, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, chn9.inner.ctx.Err()
	case <-chn10.inner.ctx.Done():
		return RECV_10, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, chn10.inner.ctx.Err()
	}
}
