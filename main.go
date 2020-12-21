package main

import "migomeal/rest"

func main(){
	adapter := rest.NewAdapter()
	adapter.ListenAndServe()
}
