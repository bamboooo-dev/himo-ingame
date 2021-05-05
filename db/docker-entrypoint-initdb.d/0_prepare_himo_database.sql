CREATE DATABASE IF NOT EXISTS himo;

USE himo;

DROP TABLE IF EXISTS themes;

CREATE TABLE themes
(
  id           BIGINT NOT NULL PRIMARY KEY,
  sentence     VARCHAR(100) NOT NULL
);

INSERT INTO themes (id, sentence) VALUES
(1, '無人島に持っていきたいもの'),
(2, '親になって欲しいキャラ'),
(3, '強そうな効果音'),
(4, 'グッとくる仕草'),
(5, 'キャンプに欠かせないもの'),
(6, '言われて傷つく異性からの一言'),
(7, '一人暮らしをする時に欲しいもの'),
(8, '雪山に遭難した時に欲しいもの'),
(9, '異性にモテる部活'),
(10, '異性にモテる趣味'),
(11, 'されたいプロポーズの言葉'),
(12, '小学生の時にモテる要素'),
(13, '合コンでウケる職業'),
(14, '死ぬ直前にしたいこと'),
(15, '死ぬ直前に食べたいもの'),
(16, '人気のスナック菓子'),
(17, 'カラオケの人気曲'),
(18, '母親の手料理といえば'),
(19, '付き合ってみたい芸能人'),
(20, '印象に残っているテレビ番組'),
(21, '友達の家で出てきたら嬉しいお菓子'),
(22, '大学生に人気のアルバイト'),
(23, '1日で100万円を使い切るならすること'),
(24, '人気の結婚相手の国籍'),
(25, '透明人間になったらしたいこと'),
(26, '正月に食べるもの'),
(27, '100円ショップで人気の商品'),
(28, '武器になりそうな文房具'),
(29, '有名なアパレルブランド'),
(30, 'カロリーの高そうなお菓子'),
(31, '走るのが速い動物'),
(32, '人気のおにぎりの具'),
(33, '人気な寿司ネタ'),
(34, '時代遅れの言葉'),
(35, '言われて嬉しい言葉'),
(36, '上司に言われて嬉しい言葉'),
(37, '人気の告白のシチュエーション'),
(38, '人気の新婚旅行先'),
(39, '手に入れたい超能力'),
(40, 'おしゃれな街'),
(41, '流行りの技術'),
(42, '人気なプログラミング言語'),
(43, '初心者におすすめのプログラミング言語'),
(44, '人気のエディタ'),
(45, 'よく使うLinuxコマンド'),
(46, '人気のLinuxコマンド'),
(47, '技術力の高いIT企業'),
(48, 'よく使うポート番号'),
(49, '人気のフレームワーク'),
(50, 'チーム開発で大切なこと'),
(51, '有名な技術書'),
(52, '便利なGUIツール'),
(53, 'プログラミング中で楽しい瞬間'),
(54, '辛いバグ'),
(55, '人気のクラウドサービス'),
(56, '開発が盛んなOSS'),
(57, '人気のデータベース'),
(58, 'よく使うDocker image'),
(59, '人気のCIツール'),
(60, '人気のOS'),
(61, 'エンジニアに人気の趣味'),
(62, '人気のライブラリ'),
(63, '強いエンジニアの共通点'),
(64, '強いエンジニア'),
(65, 'よく使うgitコマンド'),
(66, 'エンジニアに人気のWikiツール'),
(67, 'エンジニアに人気のゲーム'),
(68, 'エンジニアに多い出身学部・学科'),
(69, 'エンジニアの良くないところ'),
(70, 'エンジニアの良いところ'),
(71, 'エンジニアの好きな飲み物'),
(72, 'エンジニアが恋人に求めるであろうポイント'),
(73, 'すぐに伸びそうな人の特徴'),
(74, 'エンジニアがよく着る服'),
(75, 'これから来そうな言語'),
(76, 'これから来そうなサービス'),
(77, 'これから来そうな技術'),
(78, '先輩エンジニアに言われて嬉しい言葉'),
(79, 'エンジニアのデスクに必ずあるもの'),
(80, '初心者の頃やりがちなミス'),
(81, 'お持ち帰りの誘い文句'),
(82, 'チラ見えしたらドキッとする部位'),
(83, '行為中に興奮する異性からの言葉'),
(84, '童貞がイチコロになる服装'),
(85, '初めての行為でやりがちなミス'),
(86, '人気なAV女優といえば'),
(87, 'ピロートークの一言目といえば'),
(88, '風俗あるある'),
(89, 'エロい女にありがちなLINEのトプ画'),
(90, 'ヤリモク男の特徴'),
(91, '実はエロい女の特徴'),
(92, '経験人数が少なさそうな男の特徴'),
(93, '行為が上手い男の特徴'),
(94, 'なんかエロい擬音語'),
(95, '女遊びする男の部屋の特徴'),
(96, '相手にしてもらいたいコスプレ'),
(97, '手に入れたいエロい超能力'),
(98, '浮気癖のある人の特徴'),
(99, '日常生活で男性にエロいと感じる瞬間'),
(100, '日常生活で女性にエロいと感じる瞬間'),
(101, 'すぐに彼氏が入れ替わる女の特徴'),
(102, 'セフレに重視する点'),
(103, 'ゆるふわビッチが好きなブランド'),
(104, '実は性欲が強い男の特徴'),
(105, 'バレたくない性癖'),
(106, '抱かれたい男性芸能人'),
(107, '抱きたい女性芸能人'),
(108, '好きなAVのカテゴリー'),
(109, 'おもちゃに使えそうな食べ物'),
(110, '一夜を共にしたい相手の国籍'),
(111, '興奮するシチュエーション'),
(112, '男性に人気のカップ数'),
(113, 'エロい人が多い部活'),
(114, '人気のフェチ'),
(115, 'セクシーと感じる仕草'),
(116, '行為に至るまでの理想のシチュエーション'),
(117, 'ムラムラするとき'),
(118, '一度はやってみたい場所'),
(119, 'それまで満々だったヤる気が失せる時'),
(120, '行為中に幸せを感じるとき');
