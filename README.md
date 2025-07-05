# Galeria-Dinamica-de-Personas-Filtra-Busca-y-Marca-Favoritos-Go-Vue.js-MySQL

Galería de personas en una interfaz moderna y rápida. Filtra por género (hombres y mujeres), busca perfiles por nombre y marca tus favoritos con un simple 'like' persistente. Desarrollado con Go para el backend, MySQL para la base de datos y Vue.js en el frontend, ofrece una experiencia de usuario fluida y reactiva sin recargar la página.

Una aplicación web moderna y reactiva diseñada para explorar una galería de perfiles de personas. Los usuarios pueden **filtrar por género**, **buscar personas por nombre** y **marcar sus perfiles favoritos con un 'Me gusta'** persistente, todo ello sin recargar la página.

Este proyecto ha sido desarrollado siguiendo las mejores prácticas de **UI/UX**, **SEO**, **seguridad** y **accesibilidad**, y está optimizado para ofrecer un rendimiento excelente.

## Características Destacadas

* **Galería Interactiva:** Visualiza una amplia colección de perfiles de personas (hombres y mujeres).
* **Filtros por Género:** Filtra instantáneamente la galería para mostrar solo personas de género masculino, femenino o todos.
* **Búsqueda por Nombre:** Encuentra rápidamente perfiles específicos utilizando la función de búsqueda por nombre.
* **Sistema de 'Me Gusta' Persistente:** Haz clic en el icono de corazón para marcar tus personas favoritas. El estado del 'Me gusta' se guarda en la base de datos y se mantiene al recargar la página o el servidor.
* **Experiencia de Usuario Fluida:** Todas las interacciones (filtrado, búsqueda, likes) ocurren sin necesidad de recargar la página, gracias a una comunicación eficiente entre el frontend y el backend.
* **Diseño Responsivo:** La galería se adapta automáticamente a diferentes tamaños de pantalla, ofreciendo una experiencia óptima en dispositivos móviles y de escritorio.
* **Optimización de Imagen:** Implementación de lazy loading (`loading="lazy"`) para una carga de imágenes eficiente y `object-fit: cover` para una visualización consistente.
* **SEO & Compartibilidad:** Meta etiquetas Open Graph y Twitter Cards configuradas para una previsualización atractiva al compartir enlaces en redes sociales.
* **Validaciones:** Integración de validaciones en el backend para asegurar la integridad de los datos.

## Tecnologías Utilizadas

* **Backend:** [**Go (Golang)**](https://go.dev/)
    * Framework Web: [Gin Gonic](https://gin-gonic.com/)
    * Controladores RESTful para la gestión de personas y likes.
    * Conexión eficiente con la base de datos.
* **Base de Datos:** [**MySQL**](https://www.mysql.com/)
    * Almacenamiento de perfiles de personas y el estado de 'Me gusta' de los usuarios.
    * Uso de índices para optimizar las consultas de filtrado y búsqueda.
* **Frontend:** [**Vue.js**](https://vuejs.org/)
    * JavaScript Framework progresivo para construir interfaces de usuario interactivas.
    * Componentes reactivos para una experiencia de usuario dinámica.
    * Axios (o Fetch API) para las solicitudes HTTP asíncronas.
* **Manejo de Bases de Datos:** `database/sql` para Go, `go-sql-driver/mysql`.
* **Estilos:** CSS puro con `scoped styles` en Vue para encapsulación de componentes y `CSS Grid` para un layout responsivo.

## Configuración y Ejecución Local

Sigue estos pasos para levantar la aplicación en tu entorno local.

### Prerequisitos

* [Go (Golang)](https://go.dev/doc/install) (versión 1.18 o superior)
* [Node.js](https://nodejs.org/en/download/) y [npm](https://www.npmjs.com/get-npm)
* [MySQL Server](https://dev.mysql.com/downloads/mysql/)

### 1. Configuración de la Base de Datos

Asegúrate de que tu servidor MySQL esté corriendo.

```sql
-- Conéctate a tu MySQL CLI o usa un cliente como MySQL Workbench
-- Crea la base de datos si no existe
CREATE DATABASE IF NOT EXISTS hot_or_not_db;

-- Usa la base de datos
USE hot_or_not_db;

-- Crea la tabla 'personas'
CREATE TABLE IF NOT EXISTS personas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    edad INT,
    genero VARCHAR(50),
    url_imagen VARCHAR(2048) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Crea la tabla 'users' (necesaria para la clave foránea de likes)
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inserta un usuario dummy para pruebas (siempre usa el ID 1 para los likes)
INSERT IGNORE INTO users (id, username) VALUES (1, 'test_user');

-- Crea la tabla 'user_likes' para almacenar los likes de los usuarios
CREATE TABLE IF NOT EXISTS user_likes (
    user_id INT NOT NULL,
    persona_id INT NOT NULL,
    PRIMARY KEY (user_id, persona_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (persona_id) REFERENCES personas(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- (Opcional) Inserta datos de ejemplo en la tabla 'personas'
Ejemplo:
INSERT INTO personas (nombre, edad, genero, url_imagen) VALUES
('Felipe Bravo', 33, 'MASCULINO', 'https://images.unsplash.com/photo-1531891437562-4301efdf8aad?auto=format&fit=crop&q=80&w=1974&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Valeria Soto', 25, 'FEMENINO', 'https://images.unsplash.com/photo-1529626465034-45b428ab42e0?auto=format&fit=crop&q=80&w=1974&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Ricardo Núñez', 40, 'MASCULINO', 'https://images.pexels.com/photos/10892019/pexels-photo-10892019.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2'),
('Brenda Silva', 29, 'FEMENINO', 'https://images.pexels.com/photos/1572808/pexels-photo-1572808.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2')

Configuración y Ejecución del Backend (Go)

Navega a la carpeta backend:

cd backend

Instala las dependencias de Go:

go mod tidy

Configura la variable de entorno MYSQL_DSN. Esta variable debe contener la cadena de conexión a tu base de datos MySQL.

En Linux/macOS (Bash/Zsh) o Git Bash (Windows):

export MYSQL_DSN="root:@tcp(127.0.0.1:3306)/hot_or_not_db?charset=utf8mb4&parseTime=True&loc=Local"
export FRONTEND_URL="http://localhost:5173" # Para CORS

(Asegúrate de reemplazar root:@ con tus credenciales si no usas root sin contraseña).

Ejecuta el servidor Go:

go run main.go

Verás un mensaje indicando que el servidor está escuchando en el puerto 8000.

Configuración y Ejecución del Frontend (Vue.js)

Abre una nueva ventana de terminal y navega a la carpeta frontend:

cd frontend

Instala las dependencias de Node.js:

Instala las dependencias de Node.js:

Inicia el servidor de desarrollo de Vue:

npm run dev

El servidor frontend se iniciará, generalmente en http://localhost:5173/.

¡Listo!
Abre tu navegador y visita http://localhost:5173/ para ver la Galería Dinámica de Personas en acción.

Contribución
¡Las contribuciones son bienvenidas! Siéntete libre de abrir un issue o enviar un pull request.







