# scrape-vsezaodvoz

Scrape data from https://vsezaodvoz.cz/ website.


## Roadmap

Things to do:

- Use paging (visist page 1..N).
- Flags:
  - region
  - type
  - new
  - with_photo
  - no_reservation
  - district  

## Build

```powershell
$env:GOOS="windows"
$env:GOARCH="amd64"
go build
```

## Usage

### Examples

```powershell
.\scrape-vsezaodvoz.exe
.\scrape-vsezaodvoz.exe > output.txt
.\scrape-vsezaodvoz.exe -c dum-a-zahrada > output.txt
.\scrape-vsezaodvoz.exe -c dum-a-zahrada/kychyn > output.txt
```

```
-c <name> Select category 
-l <1|0>  Should limit to 24h old?
```

The result is redirected to standard output (`stdout`) and each line is a valid JSON object. See the [example](./output.txt). Without redirection there is a mix of result and logging messages (`stderr`) in console.

### Categories

Select category from the list bellow e.g `dum-a-zahrada` or with subcategory `dum-a-zahrada/kuchyn`

- Auto-moto (1 / auto-moto)
  - Auto
  - Moto  
  - Ostatní (1.3 / ostatni)
- Dům a zahrada (2 /dum-a-zahrada)
  - Kuchyň (2.1 / kuchyn)
  - Nábytek (2.2 / nabytek)
  - Nářadí (2.3)
  - Ostatní (2.4)
  - Stavebniny (2.5)
  - Vybavení (2.6)
  - Zahrada (2.7)
- Elektro (3)
  - Audio (3.1)
  - Foto (3.2)
  - Malé spotřebiče (3.3)
- Ostatní (4)
  - Počítače (4.1)
  - Velké spotřebiče (4.2)
  - Video (4.3)
- Hudba a film (5)
  - Aparatura (5.1)
  - Film (5.2)
  - Hudba (5.3)
  - Hudební nástroje (5.4)
  - Ostatní (5.5)
- Kanceláře a provozovny (6)
- Knihy a tiskoviny (7)
- Kosmetika (8)
- Oblečení (9)
- Ostatní (10)
- Potraviny (11)
- Pro děti (12)
- Sport a cestování (13)
- Umění a sběratelství (14)
- Zvířata a chovatelství (15)

(GitHub numbered lists rendering are actually messy!)
