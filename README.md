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
  - Auto
  - Moto
  - Ostatní
2. Dům a zahrada
  - Kuchyň
  - Nábytek
  - Nářadí
  - Ostatní
  - Stavebniny
  - Vybavení
  - Zahrada
3. Elektro
  - Audio
  - Foto
  - Malé spotřebiče
  - Ostatní
  - Počítače
  - Velké spotřebiče
  - Video
4. Hudba a film
  - Aparatura
  - Film
  - Hudba
  - Hudební nástroje
  - Ostatní
- (todo following)
5. Kanceláře a provozovny
6. Knihy a tiskoviny
7. Kosmetika
8. Oblečení
9. Ostatní
10. Potraviny
11. Pro děti
12. Sport a cestování
13. Umění a sběratelství
14. Zvířata a chovatelství
