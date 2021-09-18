#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Desc   :

class EnsembleBase:
    def __init__(self, max_trial_num, metric):
        """
        模型融合基类

        Parameters
        ----------
        max_trial_num : int
            使用的最大 trial 数量
        metric : str
            需要提升效果的指标
        """
        self.max_trial_num = max_trial_num
        self.metric = metric
        self.indices = None

    def fit(self, val_preds, val_label):
        """
        训练融合算法，基于所有 trial 的预测结果及标签
        返回需要用到 trial_id 列表

        Parameters
        ----------
        val_preds : list of np.ndarray
            每个 trial 在验证集上的预测结果，可以是把训练集分成两份(hold_out)或交叉验证(cv)
            顺序应该与 trial_id 保持一致
        val_label : np.ndarray
            验证集标签

        Returns
        ----------
        list of trial_id
        """
        pass

    def predict(self, test_preds):
        """
        融合预测结果

        Parameters
        ----------
        test_preds : list of np.ndarray
            用于融合的 trial 在测试集上的预测结果
            顺序应该与 fit 阶段返回结果保持一致
        val_label : np.ndarray
            测试集标签

        Returns
        ----------
        np.ndarray
        """

