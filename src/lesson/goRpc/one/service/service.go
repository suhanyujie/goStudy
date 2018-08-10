package service

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arich int

// todo
func (t *Arich) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}


