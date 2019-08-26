SRC = cmd/zip-file/*.go
OUT = zip-file


CC = go
go = $(shell which go 2> /dev/null)

ifeq (, $(go))
	echo "Go not found!"
endif

$(OUT):clean $(SRC)
	@printf "\e[33mBuilding\e[90m %s\e[0m\n" $@
	@$(CC) build -o $(OUT) $(SRC)
	@chmod 777 $(OUT)
	@printf "\e[34mDone!\e[0m\n"	
clean:
	@rm -f $(OUT)
	@printf "\e[34mAll clear!\e[0m\n"

