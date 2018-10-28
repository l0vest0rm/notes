let fs = require('fs'),
PDFParser = require("pdf2json");

let pdfParser = new PDFParser();

pdfParser.on("pdfParser_dataError", errData => console.error(errData.parserError) );
pdfParser.on("pdfParser_dataReady", pdfData => {
    fs.writeFile("./resume.json", JSON.stringify(pdfData), (err) => {
        if(err) console.log('error', err);
      });
});

pdfParser.loadPDF("./resume.pdf");