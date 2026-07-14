import os
import time
import shutil

from playwright.sync_api import sync_playwright

FRONTEND = "D:/WorkBuddy/InjectiveGlobalCup/frontend/index.html"
BACKEND = "http://localhost:8080"
OUT = "D:/WorkBuddy/InjectiveGlobalCup/demos/screenshots"
os.makedirs(OUT, exist_ok=True)

# 清理旧截图
for f in os.listdir(OUT):
    if f.endswith(".png"):
        os.remove(os.path.join(OUT, f))


def shoot(page, name, wait=2500):
    path = os.path.join(OUT, name)
    page.screenshot(path=path, full_page=False)
    print(f"  saved {name}")
    time.sleep(0.3)


with sync_playwright() as p:
    browser = p.chromium.launch(args=["--no-sandbox"])
    page = browser.new_page(viewport={"width": 1280, "height": 900})
    page.goto("file://" + FRONTEND)
    page.wait_for_timeout(3500)  # 等后端数据加载

    # 1) 比赛列表 + 已完赛比分
    shoot(page, "01_matches.png")

    # 2) 免费预测：法国 vs 西班牙
    try:
        page.fill("#match-select", "sf-fra-esp") if page.query_selector("#match-select") else None
    except Exception:
        pass
    # 点第一个可见的预测按钮
    btns = page.query_selector_all("button")
    for b in btns:
        try:
            txt = b.inner_text()
        except Exception:
            continue
        if "免费" in txt or "Free" in txt or "预测" in txt:
            b.click()
            break
    page.wait_for_timeout(2000)
    shoot(page, "02_free_predict.png")

    # 3) 付费 AI 预测 402 流程（展示 x402 拦截）
    try:
        btns = page.query_selector_all("button")
        for b in btns:
            try:
                txt = b.inner_text()
            except Exception:
                continue
            if "付费" in txt or "AI" in txt or "Premium" in txt:
                b.click()
                break
        page.wait_for_timeout(2500)
        shoot(page, "03_x402_premium.png")
    except Exception as e:
        print("premium click skipped:", e)

    # 4) 移动端响应式一屏（展示整体）
    page.set_viewport_size({"width": 414, "height": 896})
    page.reload()
    page.wait_for_timeout(3500)
    shoot(page, "04_mobile_view.png")

    browser.close()

print("=== capture done ===")
for f in sorted(os.listdir(OUT)):
    if f.endswith(".png"):
        fp = os.path.join(OUT, f)
        print(f"{f}  {os.path.getsize(fp)//1024} KB")
