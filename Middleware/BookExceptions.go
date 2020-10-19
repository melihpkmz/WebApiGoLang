package Middleware

func TitleEmptyException()error{
	return newErrorModel(400,1000,"Title should not be emtpy")

}

func AuthorEmptyException()error{
	return newErrorModel(400,1001,"Author should not be emtpy")
}


