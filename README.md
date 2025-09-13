# hello_world: Go 学習用サンプル集

本リポジトリは、Go言語の基礎的な機能を学ぶことを目的としたサンプルコード集です。章ごとに異なるトピックを扱い、ソースコード・設定ファイル・コメント・コミット履歴からリファレンスとして価値を得られる設計を意図しています。

主な対象
- Go の基礎文法（変数宣言、型変換、配列・スライス、range、switch、ラベル付きbreak など）
- Unicode/文字列処理（UTF-8、rune、byte との変換など）
- 標準ライブラリの活用（fmt、math/cmplx、unicode/utf8 など）
- 小規模な実例を通じた設計・実装の工夫の観察

技術スタック
- 言語/フレームワーク: Go (module hello_world)
- バージョン: go 1.24.5
- ビルド/フォーマット / チェック: Makefile に fmt, vet, build ターゲット
- 依存関係: 標準ライブラリのみ使用（外部依存は無し）
- 実行対象ファイル群: chapter1 〜 chapter4 のサブディレクトリ

主な機能一覧
- chapter1/hello.go
  - Hello の文字列出力の基本例。現状のコードは fmt.Println 内部でフォーマット文字列を出力しており、実行時にはフォーマット文字列の適用が行われない点が観察されます（修正候補）。
- chapter2/ex0201.go
  - complex(2.5, 3.1) など複素数の演算（加減乗除）と real/imag の取得、cmplx.Abs の計算、型変換の練習、calc1/calc2/calc3 による型変換・定数の扱いの例を含む。
- chapter3/ex0303.go
  - スライスの作成・キャパシティ・サブスライスの挙動、utf8.RuneCountInString の活用、rune/文字列の変換、switch/case の利用、ラベル付きbreak のデモ。
- chapter4/ex0422.go
  - Unicode 文字列の処理、RuneCountInString の活用、各種 switch の活用、ラベル付きbreak のデモ、長い文字列の処理などを含む。

設計・実装の工夫
- スライスの容量とサブスライスの安全性を教育的に示すサンプル（chapter3/ex0303.go の calc2）
-UTF-8/Unicode の理解を促すデモ（chapter3/ex0303.go、chapter4/ex0422.go）
- 文字列と rune の違いを実践的に扱う例（chapter2/ex0201.go、chapter4/ex0422.go）
- コードの可読性を高めるためのコメント・出力メッセージ設計

セットアップ & 動作確認方法
- ローカル環境
  - go version が 1.24.5 以上であることを確認
  - go.mod の module 名は hello_world
  - Makefile のターゲット
    - make fmt: go fmt ./...
    - make vet: go vet ./...
    - make build: vet を経由して go build
- 実行例
  - chapter1 の実行: go run chapter1/hello.go
  - chapter2 の実行: go run chapter2/ex0201.go
  - chapter3 の実行: go run chapter3/ex0303.go
  - chapter4 の実行: go run chapter4/ex0422.go
- 動作確認のコツ
  - 章ごとに出力が異なるため、期待出力をコードのコメントと照らし合わせて理解を深める
  - chapter1/hello.go の現状は fmt.Println でフォーマット文字列を出力しているが、実際には fmt.Printf の利用が適切な箇所があるため、改修の余地がある点を確認する

改善ポイント / TODO
- テスト
  - 各章の主要機能をユニットテストで検証する導入を検討
  - chapter1/hello.go の出力フォーマットの不整合を修正する前提のテスト追加
- エラーハンドリング
  - 現状サンプルではエラーハンドリングが最小限。実運用に近づける場合は入力検証等の追加を検討
- ドキュメント
  - チャプターごとの意図・使い方・出力例を README の各セクションに整理
- CI/CD
  - lint ジョブの追加、Go のビルド・テストの自動実行、PR の自動化などを追加検討
- その他
  - chapter1/hello.go の出力の修正（fmt.Printf への変更など）を検討
  - 章間の共通パターン・反復コードの抽象化の検討

強調したいポイント
- Go の基礎文法と標準ライブラリの組み合わせを学習できる点
- UTF-8/Unicode の扱いと、実践的な文字列操作の理解を深められる点
- 小規模ながら実用的なコードの設計・実装の工夫を観察できる点

Git 操作 (提案ワークフロー)
- ブランチ作成と README の更新
  - git fetch origin
  - git checkout -B feature/readme origin/main || git checkout -b feature/readme
  - README.md を生成/更新
  - git add README.md
  - git commit -m "docs: update README with overview, stack, setup, and categorized TODOs"
  - git push -u origin feature/readme
  - プルリクエストを作成して main へマージ（ブランチは残す）

次のステップ
- この README.md をリポジトリに追加しました。追加内容を確認のうえ、次の Git 操作（ブランチ作成・コミット・プッシュ・PR作成）を進めますか？
- README の内容をさらに細分化したい、もしくは特定のセクションを優先して強化したい場合は指示をください。

- [ ] 完了前提: README.md の内容を人力で確認・微調整
- [ ] 次のアクション: Git ブランチ作成と README のコミット・プッシュ
- [ ] 最終確認: PR 作成・マージ対応
