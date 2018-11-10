VERSION=$(shell git describe --tags)
NOMBRE=buchonip

N=[0m
R=[00;31m
G=[01;32m
Y=[01;33m
B=[01;34m
L=[01;30m

comandos:
	@echo ""
	@echo "${B}Comandos disponibles para ${G}${NOMBRE}${N} (versión: ${VERSION})"
	@echo ""
	@echo "    ${G}ejecutar${N}                    Pone la aplicación en funcionamiento."
	@echo "    ${G}test${N}                        Ejecuta los tests."
	@echo "    ${G}version${N}                     Publica una nueva versión."
	@echo ""


test:
	go test

ejecutar:
	go run buchonip.go

deploy:
	@echo "${G}Compilando ...${N}"
	env GOOS=linux GOARCH=amd64 go build
	@echo "${G}Copiando binario a la vps ...${N}"
	scp buchonip buchonip@hugoruscitti.com.ar:/home/buchonip
	@echo "${G}Reiniciando servicio en la vps ...${N}"
	ssh -t buchonip@hugoruscitti.com.ar "sudo /bin/systemctl restart buchonip"

version:
	./bumpversion.sh
