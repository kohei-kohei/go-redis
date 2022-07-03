USE go_database;

DROP TABLE IF EXISTS bread;

CREATE TABLE bread (
  id         INT(5) PRIMARY KEY AUTO_INCREMENT NOT NULL,
  name       VARCHAR(20) NOT NULL UNIQUE,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO bread (name) VALUES ("食パン");
INSERT INTO bread (name) VALUES ("クロワッサン");
INSERT INTO bread (name) VALUES ("メロンパン");
INSERT INTO bread (name) VALUES ("カレーパン");
INSERT INTO bread (name) VALUES ("あんぱん");
INSERT INTO bread (name) VALUES ("カツサンド");
INSERT INTO bread (name) VALUES ("ウインナーパン");
INSERT INTO bread (name) VALUES ("フランスパン");
