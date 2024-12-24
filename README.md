## Esoft
To run this project:

1. check for the ```.env```(you can copy all lines from ```.env.example``` and paste them in ```.env```) file.
2. Download from https://drive.google.com/drive/folders/1kbDHoadxANPUzszRv0cmzbTy3eHAdVAp and unpack to ```.csv``` dir. Example tree of directory:

```
csv
├── Session 1
│   ├── agents.csv
│   └── clients.csv
├── Session 2
│   ├── apartments.csv
│   ├── districts.csv
│   ├── houses.csv
│   └── lands.csv
├── Session 3
│   ├── apartment-demands.csv
│   ├── house-demands.csv
│   ├── land-demands.csv
│   └── supplies.csv
├── Session 4
│   └── deals.csv
├── Session 5
│   └── Mobile API specifications.docx
└── Session 6
```

3. Write in CLI:

```
make up
```

And we have swagger in http://localhost:8080/docs/index.html
