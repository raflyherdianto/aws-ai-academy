#!/usr/bin/env python3
"""
Fetch semua data wilayah (provinces + regencies) dari wilayah.id
dan simpan ke backend/data/wilayah.json untuk digunakan saat seeding DB.

Usage:
    python scripts/fetch_wilayah.py
"""

import json
import time
import urllib.request
import urllib.error
from pathlib import Path

BASE_URL = "https://wilayah.id/api"
OUTPUT_PATH = Path(__file__).parent.parent / "backend" / "data" / "wilayah.json"


def fetch_json(url: str, retries: int = 3) -> dict:
    for attempt in range(retries):
        try:
            with urllib.request.urlopen(url, timeout=15) as resp:
                return json.loads(resp.read().decode())
        except Exception as e:
            if attempt < retries - 1:
                print(f"  Retry {attempt + 1}/{retries} for {url}: {e}")
                time.sleep(1.5)
            else:
                raise


def main():
    OUTPUT_PATH.parent.mkdir(parents=True, exist_ok=True)

    print("[Fetch] provinces...")
    provinces_data = fetch_json(f"{BASE_URL}/provinces.json")
    provinces = provinces_data.get("data", [])
    print(f"  OK: {len(provinces)} provinces")

    all_regencies = []
    failed = []

    for i, prov in enumerate(provinces, 1):
        code = prov["code"]
        name = prov["name"]
        try:
            reg_data = fetch_json(f"{BASE_URL}/regencies/{code}.json")
            regs = reg_data.get("data", [])
            for r in regs:
                all_regencies.append({
                    "code": r["code"],
                    "province_code": code,
                    "name": r["name"]
                })
            print(f"  [{i:2}/{len(provinces)}] OK  {name} ({len(regs)} kota/kab)")
        except Exception as e:
            print(f"  [{i:2}/{len(provinces)}] ERR {name}: {e}")
            failed.append(code)
        # Gentle rate limit
        time.sleep(0.1)

    result = {
        "provinces": provinces,
        "regencies": all_regencies
    }

    with open(OUTPUT_PATH, "w", encoding="utf-8") as f:
        json.dump(result, f, ensure_ascii=False, separators=(",", ":"))

    size_kb = OUTPUT_PATH.stat().st_size / 1024
    print(f"\nDone!")
    print(f"  Provinces : {len(provinces)}")
    print(f"  Regencies : {len(all_regencies)}")
    print(f"  File      : {OUTPUT_PATH} ({size_kb:.1f} KB)")
    if failed:
        print(f"  WARN failed: {failed}")


if __name__ == "__main__":
    main()
