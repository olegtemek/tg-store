CREATE TABLE Files (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255),
  folderId INTEGER,
  query text,
  CONSTRAINT fkFolder FOREIGN KEY(folderId) REFERENCES Folders(id) ON DELETE CASCADE
)