# scrape-vsezaodvoz

Scrape data from https://vsezaodvoz.cz/ website.

## Categories

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


Example log message

```powershell
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzeraty/dum-a-zahrada?region=14&type=1&new=1&with_photo=true&no_reservation=true&district=77
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326368/rizky-kalanchoe
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326297/dozy
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326287/prirodni-kamen-ze-zahrady
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326283/matrace-140x200-pratelny-potah
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326262/propisky-pro-sberatele
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326257/zeminu
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326247/sklenene-stinidlo-k-lampe
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326240/retro-lahve-na-piti
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326239/zelezne-nohy
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326238/zelezne-hacky-i
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326237/zelezne-hacky
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326235/vaza
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326222/sklenena-misa
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326221/cibulak
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326220/plastovy-pohar-vysoky
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326219/ruzne-misky-a-poharky
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326218/vykrajovatka
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326217/vykrajvatko-na-vejce
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326197/daruji-vybaveni-bytu-vcetne-kuchynske-linky
2023/05/14 11:42:24 Visited https://vsezaodvoz.cz/inzerat/326194/kancelarske-kreslo
```
