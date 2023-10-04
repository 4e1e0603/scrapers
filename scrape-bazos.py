from playwright.async_api import async_playwright

import asyncio


"""
Bazos používá paremetr &crz={VALUE} pro offset, kde je 20 produktů na stránce.
Např. máme 54 produktů, &crz=0 zobrazí 1. až 20. produkt.
Např. máme 54 produktů, &crz=20 zobrazí 21. až 40. produkt.
Např. máme 54 produktů, &crz=40 zobrazí 40. až 54. produkt.

"""


class Collector:

    async def on(self, what, then):
        ...


async def main(keywordd):

    async with async_playwright() as p:
    
        browser = await p.chromium.launch(headless=True)
        webpage = await browser.new_page()

        async def scrape_lego(webpage, keyword):
            await webpage.goto(f'https://www.bazos.cz/search.php?hledat=lego+{keyword}+nov%C3%A9&crz=0')
            await webpage.wait_for_timeout(1000)
            
            products = await webpage.query_selector_all('.inzeraty')
            
            number_of_products_item = await webpage.query_selector(".inzeratynadpis")
            number_of_products_text = await number_of_products_item.inner_text()
            
            number_of_items = number_of_products_text.split()[-1]
            number_of_pages = int(number_of_items) // 20

            print(f"=={number_of_pages}==")
            
            for page in range(0, number_of_pages + 1): 
                print(f"--{page}--")
                await webpage.goto(f'https://www.bazos.cz/search.php?hledat=lego+{keyword}+nov%C3%A9&crz={page}')
                await webpage.wait_for_timeout(1000)
                
                products = await webpage.query_selector_all('.inzeraty')

                
                number_of_products_item = await webpage.query_selector(".inzeratynadpis")
                number_of_products_text = await number_of_products_item.inner_text()

                for product in products:
                    title = await product.query_selector('.nadpis')
                    title_text = await title.inner_text()
                    
                    price = await product.query_selector('.inzeratycena')
                    price_text = await price.inner_text()
                    link = await product.query_selector(".inzeratynadpis > a")
                    link_text = await link.get_attribute("href")
                    item = {"title": title_text, "price": price_text, "link": link_text}
                    
                    yield item
                
        async for item in scrape_lego(webpage, keyword):
            # number = "".join([x for x in title_text.split(" ") if x.isnumeric()])
            if len(item["title"]) > 0:
                print(item)

        await browser.close()

if __name__ == "__main__":

    keyword = "minecraft"

    asyncio.run(main(keyword))