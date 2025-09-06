package main

type accountResolver struct{
	server *Server
}


func (r *accountResolver) Orders(ctx context.Context,pagination *PaginationInput,query *string,i*string,obj*Account) ([]*Order,error){

}