package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open CSV file for reading
	csvFile, err := os.Open("image-data.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// Create HTML file for writing
	htmlFile, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer htmlFile.Close()

	// Write HTML file header
	htmlFile.WriteString(`<!DOCTYPE html>
<html>
<head>
	<title>My Website</title>
	<link rel="stylesheet" type="text/css" href="style.css">
</head>
<body>
	<header>
		<h1>My Website</h1>
		<p>Welcome to my website!</p>
	</header>

	<main>
		<section>
			<h2>Image Gallery</h2>
			<div class="gallery">
`)

	// Read CSV data and write HTML gallery items
	csvReader := csv.NewReader(csvFile)
	csvReader.TrimLeadingSpace = true
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// Write HTML gallery item for CSV record
		htmlFile.WriteString(fmt.Sprintf(`
				<a href="%s">
					<img src="%s" alt="%s">
					<div class="overlay">
						<h3>%s</h3>
						<p>%s</p>
					</div>
				</a>
`, record[1], record[2], record[0], record[0], record[3]))
	}

	// Write HTML file footer
	htmlFile.WriteString(`			</div>
		</section>
	</main>

	<script src="index.js"></script>
</body>
</html>
`)

	fmt.Println("Index file created successfully!")
}
