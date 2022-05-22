CXX = go
CXX_FLAGS = 
TEST_FLAGS = -count=1 -race -coverprofile=coverage.out -covermode=atomic

SRC = 
OBJ = $(SRC:.cc=.o)
EXEC = eggbite

all:

test:
	$(CXX) test $(TEST_FLAGS)

clean:
	rm -rf *.tar.gz *.out