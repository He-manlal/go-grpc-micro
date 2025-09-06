package main

type mutationResolver struct{
	server *Server
}

func (r *mutationResolver) createAccount(ctx context.context,in AccountInput) (*Account,error){
	return nil,nil
}

func (r *mutationResolver) createProduct(ctx context.context,in ProductInput) (*Product,error){
	return nil,nil

}

func (r *mutationResolver) createOrder(ctx context.context,in OrderInput) (*Order,error){
	return nil,nil

}