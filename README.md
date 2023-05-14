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
  1.1 Auto
  1.2 Moto
  1.3 Ostatní
2. Dům a zahrada
  2.1 Kuchyň
  2.2 Nábytek
  2.3 Nářadí
  2.4 Ostatní
  2.5 Stavebniny
  2.6 Vybavení
  2.7 Zahrada
3. Elektro
  3.1 Audio
  3.2 Foto
  3.3 Malé spotřebiče
  3.4 Ostatní
  3.5 Počítače
  3.6 Velké spotřebiče
  3.7 Video
4. Hudba a film
  4.1 Aparatura
  4.2 Film
  4.3 Hudba
  4.4 Hudební nástroje
  4.5 Ostatní
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
