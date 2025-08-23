# Compilador Go
GOCMD = go
GOBUILD = $(GOCMD) build
GORUN = $(GOCMD) run

# Arquivos
SEQ = sequencial.go
PAR = paralelo.go

# Alvos
all: sequencial paralelo

sequencial:
	$(GORUN) $(SEQ)

paralelo:
	$(GORUN) $(PAR)

build:
	$(GOBUILD) -o sequencial $(SEQ)
	$(GOBUILD) -o paralelo $(PAR)

clean:
	rm -f sequencial paralelo
