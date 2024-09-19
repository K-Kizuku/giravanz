import requests
from wordcloud import WordCloud
import matplotlib.pyplot as plt
from PIL import Image
import numpy as np


# ファイルの読み込み(./support_phrases.txt)

f = open('support_phrases.txt', 'r', encoding='utf-8')

text = f.read()

f.close()

# ファイルのテキストの取得
font_path = "./NotoSansJP-ExtraBold.ttf"
mask = np.array(Image.open("./giravanz3.png").convert("L"))
print(mask.dtype)

# 画像作成
wordcloud = WordCloud(width=800, height=800,   colormap="Paired",  font_path=font_path, background_color='black',mask=mask, contour_width=3, contour_color='yellow').generate(text)

# Wordcloudを表示
plt.figure(figsize=(10, 10))
plt.imshow(wordcloud)
plt.axis('off')
plt.show()

# 画像保存
wordcloud.to_file("result_wordcrowd1.png")