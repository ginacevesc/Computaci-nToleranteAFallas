# Computaci-nToleranteAFallas
Para la ejecución del proyecto se necesita instalar Go y el IDE de su preferencia, utilizamos Visual Studio Code, también se necesita tener acceso a un  navegador web para poder visualizar las páginas en HTML.
Además se necesita la instalación y configuración de Docker Compose y Docker Swarm. 
A continuación explicaremos la configuración del entorno de trabajo que realizamos en Ubuntu 20.04: 

1. Descargar Visual Studio Code:
	Desde el gestor de aplicaciones oficial de Ubuntu, buscar Visual Studio Code y descargarlo

2. Instalar Go: 
	2.1. Obtener el repositorio oficial con el comando wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz
	2.2. Descomprimir el archivo en /usr/local con el comando sudo tar -xvf go1.14.2.linux-amd64.tar.gz -C /usr/local/
	2.3. Agregar la ubicación del directorio Go a la variable de entorno $PATH con el comando export PATH=$PATH:/usr/local/go/bin
	2.4. Comprobar que Go se haya instalado correctamente corriendo el comando go version 

3. Instalar Docker:
	3.1. Instalar paquetes que permitan usar a apt usar paquetes a traves de HTTPS con el comando sudo apt install apt-transport-https ca-certificates curl software-properties-common
	3.2. Añadir la clave GPG al sistema para poder descargar la imagen de Docker de la fuente oficial con el comando curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
	3.3 Agregar el repositorio oficial de Docker con el comando sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"
	3.4. Actualizar el sistema para integrar el repositorio de Docker con sudo apt update
	3.5. Instalar Docker con el comando sudo apt install docker-ce
	3.6. Asegurar que Docker esté instalado en el sistema y checar el estado del servicio con el comando sudo systemctl status docker
	3.7. Dar permisos al servicio para evitar usar sudo con los siguientes pasos: 
		3.7.1. Agregar un nuevo grupo con el comando sudo groupadd docker
		3.7.2. Añadir nuestro usuario al grupo creado anteriormente con el comando sudo gpasswd -a $USER docker
		3.7.3. Reiniciar el equipo con sudo reboot
	3.8. Comprobar que se pueden correr las imágenes en el servicio de Docker con el comando docker run hello-world

4. Realizar Dockerfile y subir a Dockerhub
	4.1. Agregar las instrucciones a Dockerfile para ligar programa de Go con Docker:
		FROM golang:1.12.0-alpine3.9
		RUN mkdir /app
		ADD . /app
		WORKDIR /app
		RUN go build -o main .
		CMD ["/app/main"]
	4.2. Añadir las credenciales de nuestra cuenta de Dockerhub por medio del comando docker login. Si no se tiene una cuenta de Dockerhub, crearla y posteriormente correr este comando 
	4.3. Una vez creado el Dockerfile, subir a Dockerhub con el comando docker push