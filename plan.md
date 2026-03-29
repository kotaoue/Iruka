# Iruka 技術選定計画

## やりたいこと（要件）

1. Mac でデスクトップマスコットを作成する
2. 指定した画像を表示する
3. 背景色は透過する
4. ウィンドウのタイトルバー・フレームは表示しない
5. 右クリックでメニューが表示され、各種操作を依頼できる

---

## 現状の課題

[PR #7](https://github.com/kotaoue/Iruka/pull/7)・[PR #9](https://github.com/kotaoue/Iruka/pull/9) で Go + Fyne を用いた実装を試みたが、以下の問題が発生した。

- **背景透過**: Fyne のテーマ変更だけでは不十分で、CGO 経由で `NSWindow setOpaque:NO` / `setBackgroundColor:` を呼び出す必要があった。
- **ボーダーレスウィンドウ + ドラッグ移動**: `CreateSplashWindow()` でタイトルバーを除去できるが、ドラッグ移動のために再び CGO で `isMovableByWindowBackground = YES` を設定する必要があった。
- **CGO の複雑さ**: macOS 固有の機能を使うたびに CGO + Objective-C コードが増え、クロスプラットフォームビルドの管理コストが上がる。

Go + Fyne は汎用 GUI フレームワークであり、macOS ネイティブ機能との親和性が低い。今後も同様の問題が継続的に発生することが予想される。

---

## 技術選定

### 比較表

| 技術スタック | 背景透過 | ボーダーレス | 右クリックメニュー | 画像表示 | macOS ネイティブ度 | 学習コスト |
|---|---|---|---|---|---|---|
| **Go + Fyne** (現状) | △ (CGO 必要) | △ (CGO 必要) | ○ | ○ | 低 | 低 |
| **Go + macdriver** | ○ | ○ | ○ | ○ | 高 | 中 |
| **Swift + SwiftUI** | ○ | ○ | ○ | ○ | 最高 | 中 |
| **Python + tkinter** | × | △ | △ | ○ | 低 | 低 |
| **Electron** | ○ | ○ | ○ | ○ | 低 | 低～中 |

---

### 各選択肢の評価

#### Go + Fyne（現状）
- **利点**: 既存コードベースを継続利用できる。依存が少なく、Go の知識がそのまま活かせる。
- **欠点**: macOS ネイティブ機能（ウィンドウ透過・ドラッグ移動など）を使うたびに CGO + Objective-C が必要になる。CGO を使うとクロスコンパイルが困難になる。

#### Go + macdriver
- **利点**: CGO なしで Objective-C ランタイムをバインドし、macOS ネイティブ API を直接呼び出せる。Go の資産を維持しつつ macOS らしいアプリが作れる。
- **欠点**: ライブラリが比較的新しく、ドキュメントや事例が少ない。

#### Swift + SwiftUI / AppKit
- **利点**: macOS 向けのファーストパーティ言語・フレームワーク。背景透過・ボーダーレスウィンドウ・右クリックメニューはすべて数行で実現できる。Apple のサポートが手厚く、将来的な互換性も高い。
- **欠点**: Go とは別言語になるため、既存コードを流用できない。

#### Python + tkinter / PyQt
- **利点**: 実装が簡単。
- **欠点**: 背景透過の対応が不完全。配布時に Python 環境が必要。

#### Electron
- **利点**: 実績が豊富で、CSS で透過・ボーダーレスを簡単に実現できる。
- **欠点**: 実行ファイルが非常に重い（100 MB 以上）。デスクトップマスコット用途には過剰。

---

## 推奨方針

### 方針 A（推奨）: Swift + AppKit に移行する

macOS ネイティブ機能との完全な親和性を持つ Swift + AppKit を採用する。

**理由**:
- 要件（透過・ボーダーレス・右クリックメニュー・画像表示）がすべて標準 API で実現できる。
- CGO や追加ライブラリなしにシンプルなコードで実装できる。
- 今後 macOS 特有の機能（アクセシビリティ・メニューバー常駐など）を追加する際も公式 API に沿って対応できる。

```swift
// 透過ボーダーレスウィンドウの例 (AppKit)
let window = NSWindow(
    contentRect: NSRect(x: 0, y: 0, width: 100, height: 100),
    styleMask: [.borderless],
    backing: .buffered,
    defer: false
)
window.isOpaque = false
window.backgroundColor = .clear
window.level = .floating
```

### 方針 B（次善）: Go + macdriver を採用する

Go の資産を維持しつつ macOS ネイティブ API に近づける方法として [macdriver](https://github.com/progrium/macdriver) を検討する。

- Fyne の依存を除去し、macdriver で直接 `NSWindow` を操作する。
- CGO の管理コストを下げながら macOS らしい挙動を実現できる。

---

## 実装ロードマップ（方針 A: Swift 採用の場合）

- [ ] Swift パッケージとして `Iruka.xcodeproj`（または `Package.swift`）を作成
- [ ] 指定画像を透過ボーダーレスウィンドウに表示
- [ ] ウィンドウを常に最前面に表示（`.floating` レベル）
- [ ] 右クリックメニューの実装（NSMenu）
  - [ ] 表示 / 非表示の切り替え
  - [ ] 画像変更
  - [ ] 終了
- [ ] ドラッグ移動の実装
- [ ] アプリアイコン・Info.plist の設定

---

## 補足

- 初期実装は macOS 専用で問題ない（要件が macOS 限定のため）。
- 将来的にクロスプラットフォーム対応が必要になった場合は、Go + Fyne または Electron への再移行を検討する。
