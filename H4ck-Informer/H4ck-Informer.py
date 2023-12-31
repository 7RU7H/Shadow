#!/usr/bin/python3
import requests, threading, os, sys, time, subprocess, logging, argparse, re, asyncio, pkg_resources
from bs4 import BeautifulSoup 
from typing import Any, Awaitable 


class print_colors:
   HEADER = '\033[95m'
   BLUE = '\033[94m'
   CYAN = '\033[96m'
   GREEN = '\033[92m'
   WARNING = '\033[93m'
   FAIL = '\033[91m'
   ENDC = '\033[0m'
   BOLD = '\033[1m'
   UNDERLINE = '\033[4m'


async def run_sequence(*functions: Awaitable[Any]) -> None:
    for function in functions:
        await function

async def run_parallelism(*functions: Awaitable[Any]) -> None:
    await asyncio.gather(*functions)

async def get_summary(url):
   # Simulate a delay
   await asyncio.sleep(1)
   # Return a dummy summary
   return f"Summary of {url}"

async def request_url(url):
   response = requests.get(url)
   if response.status_code == 200:
       page_content = response.text
       soup = BeautifulSoup(page_content, "html.parser")
       articles = soup.find_all("tr", class_="athing")
       for article in articles:
           title = article.find("a", class_="titlelink").text
           url = article.find("a", class_="titlelink")["href"]
           print(f"{title}\n{url}\n")
           for paragraph in soup.find_all('p'):
                if re.search('TL;DR', paragraph.text, re.IGNORECASE):
                    print(paragraph.text)
                else:
                    print(f"Failed to find a TL;DR  from {url}")


urls = {
   "site1": ["url1", "url2", "url3"],
   "site2": ["url4", "url5", "url6"],
   # Add more sites as needed
}

def test_modules():
    combined = list(sys.modules.items()) + [(dist.project_name, dist.version) for dist in pkg_resources.working_set]
    for item in combined:
        try:
            if len(item) == 2:
                module_name, module = item
                print(f"{module_name}: {module.__version__}")
            else:
                project_name, version = item[0], item[1]
                print(f"{project_name}: {version}")
        except AttributeError:
            print(f"{print_colors.WARNING}{module_name} does not have a __version__ attribute {print_colors.ENDC}")

def main():
    
    # Logging
    logging.basicConfig(filename='app.log', filemode='w', format='%(name)s - %(levelname)s - %(message)s')
    logger = logging.getLogger(__name__)
    
    # Run tests
    print(print_colors.HEADER + "Testing HEADER" + print_colors.ENDC)
    print(print_colors.BLUE + "Testing BLUE" + print_colors.ENDC)
    print(print_colors.CYAN + "Testing CYAN" + print_colors.ENDC)
    print(print_colors.GREEN + "Testing GREEN" + print_colors.ENDC)
    print(print_colors.WARNING + "Testing WARNING" + print_colors.ENDC)
    print(print_colors.FAIL + "Testing FAIL" + print_colors.ENDC)
    print(print_colors.ENDC + "Testing ENDC" + print_colors.ENDC)
    print(print_colors.BOLD + "Testing BOLD" + print_colors.ENDC)
    print(print_colors.UNDERLINE + "Testing UNDERLINE" + print_colors.ENDC)

    test_modules()
    
    url = "https://news.ycombinator.com"

if __name__ == "__main__":
   main()