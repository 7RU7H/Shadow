#!/usr/bin/python3
import requests, threading, os, sys, time, subprocess, logging, argparse, re, asyncio, pkg_resources
from bs4 import BeautifulSoup 
from typing import Any, Awaitable 
from dataclasses import dataclass

# slots break in multiple inheritance AVOID, 20% efficiency over no slot dict
@dataclass(slots=True)
class Informer_Config:
    informer_path: str
    app_name: str
    # logging directory, output directory    


    def __init__(self, args):
        self.informer_path = os.exec('pwd')
        self.app_name = os.exec('')

class print_colors:
  HEADER = '\033[95m'
  OKBLUE = '\033[94m'
  OKCYAN = '\033[96m'
  OKGREEN = '\033[92m'
  WARNING = '\033[93m'
  FAIL = '\033[91m'
  ENDC = '\033[0m'
  BOLD = '\033[1m'
  UNDERLINE = '\033[4m'
  WHITE = '\033[97m'
  YELLOW = '\033[93m'
  RED = '\033[91m'
  BLACK = '\033[90m'
  MAGENTA = '\033[95m'
  GREEN = '\033[92m'
  BLUE = '\033[94m'
  CYAN = '\033[96m'
  ORANGE = '\033[33m' 
  INDIGO = '\033[34m' 
  VIOLET = '\033[35m'


async def run_sequence(*functions: Awaitable[Any]) -> None:
    for function in functions:
        await function

async def run_parallelism(*functions: Awaitable[Any]) -> None:
    await asyncio.gather(*functions)

# TODO this does nothing
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

def main():
    
    # Logging
    logging.basicConfig(filename='app.log', filemode='w', format='%(name)s - %(levelname)s - %(message)s')
    logger = logging.getLogger(__name__)
    
    print(print_colors.OKBLUE + '_________  ___  ___  _______           ___  ___  ___   ___  ________  ___  __                   ___  ________   ________ ________  ________  _____ ______   ________  ________     ' + print_colors.ENDC)
    print(print_colors.OKCYAN + '|\___   ___\\  \|\  \|\  ___ \         |\  \|\  \|\  \ |\  \|\   ____\|\  \|\  \                |\  \|\   ___  \|\  _____\\   __  \|\   __  \|\   _ \  _   \|\_____  \|\   __  \    ' + print_colors.ENDC)
    print(print_colors.OKGREEN + '\|___ \  \_\ \  \\\  \ \   __/|        \ \  \\\  \ \  \\_\  \ \  \___|\ \  \/  /|_  ____________\ \  \ \  \\ \  \ \  \__/\ \  \|\  \ \  \|\  \ \  \\\__\ \  \|____|\ /\ \  \|\  \  ' + print_colors.ENDC)
    print(print_colors.YELLOW + '     \ \  \ \ \   __  \ \  \_|/__       \ \   __  \ \______  \ \  \    \ \   ___  \|\____________\ \  \ \  \\ \  \ \   __\\ \  \\\  \ \   _  _\ \  \\|__| \  \    \|\  \ \   _  _\  ' + print_colors.ENDC)
    print(print_colors.RED + '      \ \  \ \ \  \ \  \ \  \_|\ \       \ \  \ \  \|_____|\  \ \  \____\ \  \\ \  \|____________|\ \  \ \  \\ \  \ \  \_| \ \  \\\  \ \  \\  \\ \  \    \ \  \  __\_\  \ \  \\  \| ' + print_colors.ENDC)
    print(print_colors.GREEN + '       \ \__\ \ \__\ \__\ \_______\       \ \__\ \__\     \ \__\ \_______\ \__\\ \__\              \ \__\ \__\\ \__\ \__\   \ \_______\ \__\\ _\\ \__\    \ \__\|\_______\ \__\\ _\ ' + print_colors.ENDC)
    print(print_colors.INDIGO + '        \|__|  \|__|\|__|\|_______|        \|__|\|__|      \|__|\|_______|\|__| \|__|               \|__|\|__| \|__|\|__|    \|_______|\|__|\|__|\|__|     \|__|\|_______|\|__|\|__|' + print_colors.ENDC)

     
    url = "https://news.ycombinator.com"

    parser = argparse.ArgumentParser(#prog='H4ck-Inf0rm3r',
                                    usage='%(prog)s [options] target',
                                    description='Automated InfoSec, Hacker, CVEs News Aggregator',
                                    epilog='Happy Hacking :)',
                                    add_help=True,)
    args = parser.parse_args()
    args_dict = vars(args)

    curr = Informer_Config(**args_dict)

 
    if len(sys.argv) ==  1:
        parser.print_help()
        print("")
        print("Open \`crontab -e\` and add both or adjust")
        print(f"0 5 * * * {curr.informer_path}{curr.app_name}")
        print(f"0 11 * * * {curr.informer_path}{curr.app_name}")
        print(f"0 17 * * * {curr.informer_path}{curr.app_name}")
        print(f"0 23 * * * {curr.informer_path}{curr.app_name}")
        sys.exit(1)


if __name__ == "__main__":
   main()
   sys.exit(0)