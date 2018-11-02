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
	@echo "    ${G}deploy${N}                      Actualiza la versión de producción."
	@echo ""


test:
	go test

ejecutar:
	go run buchonip.go

deploy:
	@echo "${G}Compilando ...${N}"
	env GOOS=linux GOARCH=amd64 go build
	@echo "${G}Copiando binario a la vps ...${N}"
	scp buchonip root@hugoruscitti.com.ar:/root
	@echo "${G}Reiniciando servicio en la vps ...${N}"
	ssh -t root@hugoruscitti.com.ar "systemctl restart buchonip"
