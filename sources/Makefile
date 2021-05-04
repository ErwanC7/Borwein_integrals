##
## EPITECH PROJECT, 2020
## Makefile
## File description:
## Makefile du 109titration
##

NAME	=	110borwein

SRC		=	sources/main.go \
			sources/rules.go

CC		=	go build

all :	$(NAME)

$(NAME) :
	$(CC) -o $(NAME) $(SRC)

clean :
	$(RM) $(NAME)

fclean :
	$(RM) $(NAME)

re :	fclean all

.PHONY :	all clean fclean re
