# scrape-vsezaodvoz

Scrape data from https://vsezaodvoz.cz/ website.

## Build

```powershell
$env:GOOS="windows"
$env:GOARCH="amd64"
go build
```

## Usage

```powershell
scrape-vsezaodvoz.exe
.\scrape-vsezaodvoz.exe > output.txt
```

The result is redirected to standard output (`stdout`) and each line is a valid JSON object. See the [example](./output.txt). Without redirection there is a mix of result and logging messages (`stderr`) in console.

### Categories

1. Auto-moto
    1. Auto
    2. Moto
    3. Ostatní
2. Dům a zahrada
    1. Kuchyň
    2. Nábytek
    3. Nářadí
    4. Ostatní
    5. Stavebniny
    6. Vybavení
    7. Zahrada
3. Elektro
    1. Audio
    2. Foto
    3. Malé spotřebiče
    4. Ostatní
    5. Počítače
    6. Velké spotřebiče
    7. Video
4. Hudba a film
    1. Aparatura
    2. Film
    3. Hudba
    4. Hudební nástroje
    5. Ostatní
5. Kanceláře a provozovny
6. Knihy a tiskoviny
7. Kosmetika
8. Oblečení
9. Ostatní
10. Potraviny
11. Pro děti
12. Sport a cestování
13. Umění a sběratelství
14.  Zvířata a chovatelství
