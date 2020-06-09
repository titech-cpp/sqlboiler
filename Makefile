build:
	statik -src template -f
	go build -o sqlboiler
code: ./sample/sqlboiler.yaml
	./sqlboiler code -c models -y sample/sqlboiler.yaml
schema: ./sample/sqlboiler.yaml
	./sqlboiler schema -s docs -y sample/sqlboiler.yaml
all:
	make code
	make schema