import random

# List of supportive phrases
phrases = [
    "ファイト！", "がんばれ！", "負けるな！", "いけ！", "よくやった！", "ゴールを狙え！", "勝利はすぐそこだ！", "諦めるな！",
    "君ならできる！", "チームを信じろ！", "勝てるぞ！", "ナイスプレー！", "その調子！", "突っ走れ！", "気持ちを込めて！",
    "戦え！", "勝利を信じて！", "やればできる！", "君がヒーローだ！", "最後まで走り抜け！", "スピードアップ！", 
    "チャンスをつかめ！", "強気でいこう！", "全力で行け！", "気合いだ！", "さあ、ゴールだ！", "君に期待している！", 
    "最高の試合を見せてくれ！", "勇気を出して！", "自信を持って！", "負ける気がしない！", "これが勝負だ！", 
    "さあ、一発！", "君の力を信じている！", "もう少しだ！", "ゴールまであと一歩！", "その調子で行け！", "いいぞ！", 
    "熱くなれ！", "みんなが応援している！", "ピッチを支配しろ！", "勝利を掴め！", "止まらないで！", 
    "リズムを崩さないで！", "やるしかない！", "諦めない心を見せてくれ！", "チームワークを発揮しろ！", 
    "熱いプレーを見せてくれ！", "君ならやれる！", "勝負はここからだ！","ギラヴァンツ","ギラヴァンツ北九州"
]

# Function to write repeated phrases to a text file
def save_phrases_to_file(phrases, filename):
    with open(filename, 'w', encoding='utf-8') as file:
        for phrase in phrases:
            # Randomly choose a number of repetitions for each phrase
            repeat_count = random.randint(10, 50)  # You can set this as needed
            file.write((phrase + '\n') * repeat_count)

# Save to 'support_phrases.txt'
save_phrases_to_file(phrases, 'support_phrases.txt')
print("Phrases saved to support_phrases.txt.")