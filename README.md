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

- Auto-moto
  - Auto
  - Moto
  - Ostatní
- Dům a zahrada
  - Kuchyň
  - Nábytek
  - Nářadí
  - Ostatní
  - Stavebniny
  - Vybavení
  - Zahrada
- Elektro
  - Audio
  - Foto
  - Malé spotřebiče
  - Ostatní
  - Počítače
  - Velké spotřebiče
  - Video
- Hudba a film
  - Aparatura
  - Film
  - Hudba
  - Hudební nástroje
  - Ostatní
- (todo following)
- Kanceláře a provozovny
- Knihy a tiskoviny
- Kosmetika
- Oblečení
- Ostatní
- Potraviny
- Pro děti
- Sport a cestování
- Umění a sběratelství
- Zvířata a chovatelství
