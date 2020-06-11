build:
	statik -src template -f
	go build -o sqlboiler
code: ./sample/sqlboiler.yaml
	./sqlboiler code -c sample/models -y sample/docs/sqlboiler.yaml
schema: ./sample/sqlboiler.yaml
	./sqlboiler schema -s sample/docs -y sample/docs/sqlboiler.yaml
all:
	make code
	make schema