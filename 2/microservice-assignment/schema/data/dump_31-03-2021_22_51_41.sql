
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `omdb-service` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `omdb-service`;
DROP TABLE IF EXISTS `logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `logs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uri` text,
  `response` text,
  `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `logs` WRITE;
/*!40000 ALTER TABLE `logs` DISABLE KEYS */;
INSERT INTO `logs` VALUES (1,'http://www.omdbapi.com?apikey=faf7e5bb&s=captain&p=1&y=2020&type=series','{\"Response\":\"True\",\"totalResults\":\"4\",\"Search\":[{\"imdbID\":\"tt12580610\",\"Title\":\"The Epic Tales of Captain Underpants in Space\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BY2VmMTMzNTUtZWU0Zi00Yzg4LTk1NGUtYjY2ZDk5ZDZiNDg2XkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg\"},{\"imdbID\":\"tt12604360\",\"Title\":\"Captain Fantasy\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BYjMzY2FmNzYtYWQ1MC00YjdiLTlkNzAtNDAwZjY2NzMwZWJiXkEyXkFqcGdeQXVyMjgxMzc0MDQ@._V1_SX300.jpg\"},{\"imdbID\":\"tt12715090\",\"Title\":\"Captain Manicorn\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BY2NjYTI2ZGEtODg0Yi00ZWQzLWJhYjUtZTA2Zjc2MWI1YmEyXkEyXkFqcGdeQXVyMTIxMDQzOTA1._V1_SX300.jpg\"},{\"imdbID\":\"tt12186092\",\"Title\":\"Captain Carol\'s Cosmic Conquest\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BZTEyOGRiZjEtYTFkZS00Nzc4LWEzYmQtYzI5MGVlZGQxY2NjXkEyXkFqcGdeQXVyNjgxMTIwNjY@._V1_SX300.jpg\"}]}','2021-03-31 14:22:22'),(2,'http://www.omdbapi.com?apiKey=faf7e5bb&i=tt11383906','{\"Response\":\"True\",\"imdbID\":\"tt11383906\",\"Title\":\"The 7th Rule: Captain Nog Forever\",\"Year\":\"2019–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BYjY5MzhkODQtMDBiZi00NjhlLTk0ODYtZGE0MTM0ZGQxY2Q5XkEyXkFqcGdeQXVyMTQ4NDgxMzM@._V1_SX300.jpg\",\"Rated\":\"N/A\",\"Released\":\"11 Oct 2019\",\"Runtime\":\"N/A\",\"Genre\":\"Sci-Fi\",\"Director\":\"N/A\",\"Writer\":\"N/A\",\"Actors\":\"Ryan T. Husk, Cirroc Lofton\",\"Plot\":\"N/A\",\"Language\":\"English\",\"Country\":\"USA\",\"Awards\":\"N/A\",\"Ratings\":[{\"Source\":\"Internet Movie Database\",\"Value\":\"8.7/10\"}],\"Metascore\":\"N/A\",\"imdbRating\":\"8.7\",\"imdbVotes\":\"31\"}','2021-03-31 14:38:14'),(3,'http://www.omdbapi.com?apiKey=faf7e5bb&i=tt11383906','{\"Response\":\"True\",\"imdbID\":\"tt11383906\",\"Title\":\"The 7th Rule: Captain Nog Forever\",\"Year\":\"2019–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BYjY5MzhkODQtMDBiZi00NjhlLTk0ODYtZGE0MTM0ZGQxY2Q5XkEyXkFqcGdeQXVyMTQ4NDgxMzM@._V1_SX300.jpg\",\"Rated\":\"N/A\",\"Released\":\"11 Oct 2019\",\"Runtime\":\"N/A\",\"Genre\":\"Sci-Fi\",\"Director\":\"N/A\",\"Writer\":\"N/A\",\"Actors\":\"Ryan T. Husk, Cirroc Lofton\",\"Plot\":\"N/A\",\"Language\":\"English\",\"Country\":\"USA\",\"Awards\":\"N/A\",\"Ratings\":[{\"Source\":\"Internet Movie Database\",\"Value\":\"8.7/10\"}],\"Metascore\":\"N/A\",\"imdbRating\":\"8.7\",\"imdbVotes\":\"31\"}','2021-03-31 14:39:12'),(4,'http://www.omdbapi.com?apikey=faf7e5bb&s=captain&p=1&y=2020&type=series','{\"Response\":\"True\",\"totalResults\":\"4\",\"Search\":[{\"imdbID\":\"tt12580610\",\"Title\":\"The Epic Tales of Captain Underpants in Space\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BY2VmMTMzNTUtZWU0Zi00Yzg4LTk1NGUtYjY2ZDk5ZDZiNDg2XkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg\"},{\"imdbID\":\"tt12604360\",\"Title\":\"Captain Fantasy\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BYjMzY2FmNzYtYWQ1MC00YjdiLTlkNzAtNDAwZjY2NzMwZWJiXkEyXkFqcGdeQXVyMjgxMzc0MDQ@._V1_SX300.jpg\"},{\"imdbID\":\"tt12715090\",\"Title\":\"Captain Manicorn\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BY2NjYTI2ZGEtODg0Yi00ZWQzLWJhYjUtZTA2Zjc2MWI1YmEyXkEyXkFqcGdeQXVyMTIxMDQzOTA1._V1_SX300.jpg\"},{\"imdbID\":\"tt12186092\",\"Title\":\"Captain Carol\'s Cosmic Conquest\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BZTEyOGRiZjEtYTFkZS00Nzc4LWEzYmQtYzI5MGVlZGQxY2NjXkEyXkFqcGdeQXVyNjgxMTIwNjY@._V1_SX300.jpg\"}]}','2021-03-31 15:09:13'),(5,'http://www.omdbapi.com?apiKey=faf7e5bb&i=tt11383906','{\"Response\":\"True\",\"imdbID\":\"tt11383906\",\"Title\":\"The 7th Rule: Captain Nog Forever\",\"Year\":\"2019–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BYjY5MzhkODQtMDBiZi00NjhlLTk0ODYtZGE0MTM0ZGQxY2Q5XkEyXkFqcGdeQXVyMTQ4NDgxMzM@._V1_SX300.jpg\",\"Rated\":\"N/A\",\"Released\":\"11 Oct 2019\",\"Runtime\":\"N/A\",\"Genre\":\"Sci-Fi\",\"Director\":\"N/A\",\"Writer\":\"N/A\",\"Actors\":\"Ryan T. Husk, Cirroc Lofton\",\"Plot\":\"N/A\",\"Language\":\"English\",\"Country\":\"USA\",\"Awards\":\"N/A\",\"Ratings\":[{\"Source\":\"Internet Movie Database\",\"Value\":\"8.7/10\"}],\"Metascore\":\"N/A\",\"imdbRating\":\"8.7\",\"imdbVotes\":\"31\"}','2021-03-31 15:09:15');
/*!40000 ALTER TABLE `logs` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

