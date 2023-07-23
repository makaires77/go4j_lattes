const express = require('express');
const app = express();
const path = require('path');

const PORT = process.env.PORT || 3000;

// Defina a pasta 'static' para servir os arquivos estáticos
app.use(express.static(path.join(__dirname, 'static')));

// Defina a rota para a página inicial (index.html)
app.get('/', (req, res) => {
  res.sendFile(path.join(__dirname, 'static', 'index.html'));
});

// Inicie o servidor
app.listen(PORT, () => {
    console.log(`Servidor iniciado em http://localhost:${PORT}`);
});
