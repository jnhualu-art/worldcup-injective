import os

from moviepy import ImageClip, TextClip, CompositeVideoClip, concatenate_videoclips

OUT = "D:/WorkBuddy/InjectiveGlobalCup/demos"
SHOT_DIR = os.path.join(OUT, "screenshots")
VIDEO = os.path.join(OUT, "injective_worldcup_demo.mp4")

titles = [
    ("01_matches.png", "Live 2026 World Cup knockout data"),
    ("02_free_predict.png", "Free heuristic prediction (Go backend)"),
    ("03_x402_premium.png", "x402-gated premium AI prediction"),
    ("04_mobile_view.png", "Responsive UI - works anywhere"),
]

FONT = "C:/Windows/Fonts/arial.ttf"
if not os.path.exists(FONT):
    FONT = None

clips = []
for fn, title in titles:
    p = os.path.join(SHOT_DIR, fn)
    if not os.path.exists(p):
        print("missing", fn)
        continue
    img = ImageClip(p).with_duration(4)
    img = img.resized(width=1280)
    if img.h >= 720:
        img = img.cropped(x_center=640, y_center=img.h // 2, width=1280, height=720)
    else:
        img = img.with_height(720)
    comp = img
    if FONT:
        txt = TextClip(
            text=title, font_size=34, color="white", bg_color="black",
            size=(1280, 56), font=FONT,
        ).with_duration(4).with_position(("center", "bottom"))
        comp = CompositeVideoClip([img, txt])
    clips.append(comp)

if not clips:
    raise SystemExit("no screenshots found")

final = concatenate_videoclips(clips, method="compose")
final.write_videofile(VIDEO, fps=24, codec="libx264", audio=False)
print("VIDEO_READY", VIDEO, f"{os.path.getsize(VIDEO)//1024} KB")
