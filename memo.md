# TinyLFU

[TinyLFU: A Highly Efficient Cache Admission Policy](https://arxiv.org/pdf/1512.00727.pdf)の和訳

## Abstract

この論文では、偏ったアクセス分布になりがちなキャッシュの効率性の増大のために、頻度ベースのcache admission policyを使うことを提案します。
新しくアクセスされた要素とキャッシュからeviction対象が与えられたら、我々の方式では、最近のアクセス履歴に基づいて
eviction対象の犠牲によってキャッシュに新しいアイテムを入れることが価値があるかどうかを決定します。

