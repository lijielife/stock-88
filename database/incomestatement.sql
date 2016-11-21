CREATE TABLE `finance`.`incomestatement` (              
	`code` varchar(6) NOT NULL,
	`reportdate` date NOT NULL,
    `biztotinco`  decimal(20,2) comment "一、营业总收入",
    `bizinco`  decimal(20,2) comment "营业收入",
    `inteinco`  decimal(20,2) comment "利息收入",
    `earnprem`  decimal(20,2) comment "已赚保费",
    `pouninco`  decimal(20,2) comment "手续费及佣金收入",
    `realsale`  decimal(20,2) comment "房地产销售收入",
    `otherbizinco`  decimal(20,2) comment "其他业务收入",
    `biztotcost`  decimal(20,2) comment "二、营业总成本",
    `bizcost`  decimal(20,2) comment "营业成本",
    `inteexpe`  decimal(20,2) comment "利息支出",
    `pounexpe`  decimal(20,2) comment "手续费及佣金支出",
    `realsalecost`  decimal(20,2) comment "房地产销售成本",
    `deveexpe`  decimal(20,2) comment "研发费用",
    `surrgold`  decimal(20,2) comment "退保金",
    `compnetexpe`  decimal(20,2) comment "赔付支出净额",
    `contress`  decimal(20,2) comment "提取保险合同准备金净额",
    `polidiviexpe`  decimal(20,2) comment "保单红利支出",
    `reinexpe`  decimal(20,2) comment "分保费用",
    `otherbizcost`  decimal(20,2) comment "其他业务成本",
    `biztax`  decimal(20,2) comment "营业税金及附加",
    `salesexpe`  decimal(20,2) comment "销售费用",
    `manaexpe`  decimal(20,2) comment "管理费用",
    `finexpe`  decimal(20,2) comment "财务费用",
    `asseimpaloss`  decimal(20,2) comment "资产减值损失",
    `valuechgloss`  decimal(20,2) comment "公允价值变动收益",
    `inveinco`  decimal(20,2) comment "投资收益",
    `assoinveprof`  decimal(20,2) comment "其中:对联营企业和合营企业的投资收益",
    `exchggain`  decimal(20,2) comment "汇兑收益",
    `futuloss`  decimal(20,2) comment "期货损益",
    `custinco`  decimal(20,2) comment "托管收益",
    `subsidyincome`  decimal(20,2) comment "补贴收入",
    `otherbizprof`  decimal(20,2) comment "其他业务利润",
    `perprofit`  decimal(20,2) comment "三、营业利润",
    `nonoreve`  decimal(20,2) comment "营业外收入",
    `nonoexpe`  decimal(20,2) comment "营业外支出",
    `noncassetsdisl`  decimal(20,2) comment "非流动资产处置损失",
    `totprofit`  decimal(20,2) comment "利润总额",
    `incotaxexpe`  decimal(20,2) comment "所得税费用",
    `unreinveloss`  decimal(20,2) comment "未确认投资损失",
    `netprofit`  decimal(20,2) comment "四、净利润",
    `parenetp`  decimal(20,2) comment "归属于母公司所有者的净利润",
    `mergeformnetprof`  decimal(20,2) comment "被合并方在合并前实现净利润",
    `minysharrigh`  decimal(20,2) comment "少数股东损益",
    `basiceps`  decimal(20,2) comment "五、基本每股收益",
    `dilutedeps`  decimal(20,2) comment "稀释每股收益",
    `othercompinco`  decimal(20,2) comment "六、其他综合收益",
    `parecompinco`  decimal(20,2) comment "归属于母公司所有者的其他综合收益",
    `minysharinco`  decimal(20,2) comment "归属于少数股东的其他综合收益",
    `compincoamt`  decimal(20,2) comment "七、综合收益总额",
    `parecompincoamt`  decimal(20,2) comment "归属于母公司所有者的综合收益总额",
    `minysharincoamt`  decimal(20,2) comment "归属于少数股东的综合收益总额",
    `earlyundiprof`  decimal(20,2) comment "年初未分配利润",
    `undisprofredu`  decimal(20,2) comment "减少注册资本减少的未分配利润",
    `otherinto`  decimal(20,2) comment "其他转入",
    `otherdistprof`  decimal(20,2) comment "可分配利润",
    `legalsurp`  decimal(20,2) comment "提取法定盈余公积",
    `welfare`  decimal(20,2) comment "提取公益金",
    `capitalreserve`  decimal(20,2) comment "提取资本公积金",
    `staffaward`  decimal(20,2) comment "提取职工奖福基金",
    `reservefund`  decimal(20,2) comment "提取储备基金",
    `developfund`  decimal(20,2) comment "提取企业发展基金",
    `profreturninvest`  decimal(20,2) comment "利润归还投资",
    `supplycurcap`  decimal(20,2) comment "补充流动资本",
    `avaidistshareprof`  decimal(20,2) comment "可供股东分配的利润",
    `preferredskdiv`  decimal(20,2) comment "应付优先股股利",
    `freesurplu`  decimal(20,2) comment "提取任意公积",
    `dealwithdivi`  decimal(20,2) comment "应付普通股股利",
    `capitalizeddivi`  decimal(20,2) comment "转作资本股本的普通股股利",
    `undisprofit`  decimal(20,2) comment "未分配利润",
    `selldepartgain`  decimal(20,2) comment "出售处置部门或被投资单位所得收益",
    `natudisasterloss`  decimal(20,2) comment "自然灾害发生的损失",
    `accpolicychg`  decimal(20,2) comment "会计政策变更增加减少利润总额",
    `accestimatechg`  decimal(20,2) comment "会计估计变更增加减少利润总额",
    `debtrestruloss`  decimal(20,2) comment "债务重组损失",
    `othersupply`  decimal(20,2) comment "其他补充资料",
    PRIMARY KEY (`code`,`reportdate`)
) COMMENT='利润表';



                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                
                