# scrape-vsezaodvoz

Scrape data from <https://vsezaodvoz.cz/> website.

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

The result is redirected to standard output (`stdout`) and each line is a valid JSON objec<li>See the [example](./output.txt). Without redirection there is a </li>mix of result and logging messages (`stderr`) in console.

### Categories

<ol>
    <li>Auto-moto</li>
    <ol>
        <li>Auto</li>
        <li>Moto</li>
        <li>Ostatní</li>
    </ol>
    <li>Dům a zahrada</li>
    <ol>
        <li>Kuchyň</li>
        <li>Nábytek</li>
        <li>Nářadí</li>
        <li>Ostatní</li>
        <li>Stavebniny</li>
        <li>Vybavení</li>
        <li>Zahrada</li>
    </ol>
    <li>Elektro</li>
    <ol>
        <li>Audio</li>
        <li>Foto</li>
        <li>Malé spotřebiče</li>
        <li>Ostatní</li>
        <li>Počítače</li>
        <li>Velké spotřebiče</li>
        <li>Video</li>
    </ol>
    <li>Hudba a film</li>
    <ol>
        <li>Aparatura</li>
        <li>Film</li>
        <li>Hudba</li>
        <li>Hudební nástroje</li>
        <li>Ostatní</li>
    </ol>
    <li>Kanceláře a provozovny</li>
    <li>Knihy a tiskoviny</li>
    <li>Kosmetika</li>
    <li>Oblečení</li>
    <li>Ostatní</li>
    <li>Potraviny</li>
    <li>Pro děti</li>
    <li>Sport a cestování</li>
    <li>Umění a sběratelství</li>
    <li> Zvířata a chovatelství</li>
</ol>
