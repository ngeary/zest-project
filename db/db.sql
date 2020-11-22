create table applications (
    row_id varchar(55) not null,
    id int,
    member_id varchar(55),
    first_name varchar(55),
	last_name varchar(55),
	`address` varchar(55),
	dob varchar(55),
	phone varchar(55),
	loan_amnt int,
	funded_amnt int,
	funded_amnt_inv int,
	term varchar(55),
	int_rate varchar(55),
	installment int,
	grade varchar(55),
	sub_grade varchar(55),
	home_ownership varchar(55),
	annual_inc int,
	verification_status varchar(55),
	issue_d varchar(55),
	pymnt_plan varchar(55),
	`url` varchar(55),
	`desc` varchar(55),
	purpose varchar(55),
	title varchar(55),
	zip_code varchar(55),
	addr_state varchar(55),
	dti int,
	delinq_2yrs int,
	earliest_cr_line varchar(55),
	inq_last_6mths int,
	mths_since_last_delinq int,
	mths_since_last_record int,
	open_acc int,
	pub_rec int,
	revol_bal int,
	revol_util varchar(55),
	total_acc int,
	initial_list_status varchar(55),
	out_prncp int,
	out_prncp_inv int,
	total_pymnt int,
	total_pymnt_inv int,
	total_rec_prncp int,
	total_rec_int int,
	total_rec_late_fee int,
	recoveries int,
	collection_recovery_fee int,
	last_pymnt_d varchar(55),
	last_pymnt_amnt int,
	next_pymnt_d varchar(55),
	last_credit_pull_d varchar(55),
	collections_12_mths_ex_med int,
	mths_since_last_major_derog int,
	policy_code int,
	application_type varchar(55),
	annual_inc_joint int,
	dti_joint varchar(55),
	verification_status_joint int,
	acc_now_delinq int,
	tot_coll_amt int,
	tot_cur_bal int,
	open_acc_6m varchar(55),
	open_act_il varchar(55),
	open_il_12m varchar(55),
	open_il_24m varchar(55),
	mths_since_rcnt_il int,
	total_bal_il int,
	il_util int,
	open_rv_12m varchar(55),
	open_rv_24m varchar(55),
	max_bal_bc int,
	all_util varchar(55),
	total_rev_hi_lim varchar(55),
	inq_fi varchar(55),
	total_cu_tl varchar(55),
	inq_last_12m varchar(55),
	acc_open_past_24mths varchar(55),
	avg_cur_bal varchar(55),
	bc_open_to_buy varchar(55),
	bc_util varchar(55),
	chargeoff_within_12_mths int,
	delinq_amnt int,
	mo_sin_old_il_acct int,
	mo_sin_old_rev_tl_op int,
	mo_sin_rcnt_rev_tl_op int,
	mo_sin_rcnt_tl varchar(55),
	mort_acc varchar(55),
	mths_since_recent_bc int,
	mths_since_recent_bc_dlq int,
	mths_since_recent_inq int,
	mths_since_recent_revol_delinq int,
	num_accts_ever_120_pd int,
	num_actv_bc_tl int,
	num_actv_rev_tl int,
	num_bc_sats int,
	num_bc_tl int,
	num_il_tl int,
	num_op_rev_tl int,
	num_rev_accts int,
	num_rev_tl_bal_gt_0 int,
	num_sats int,
	num_tl_120dpd_2m int,
	num_tl_30dpd int,
	num_tl_90g_dpd_24m int,
	num_tl_op_past_12m int,
	pct_tl_nvr_dlq int,
	percent_bc_gt_75 int,
	pub_rec_bankruptcies int,
	tax_liens int,
	tot_hi_cred_lim int,
	total_bal_ex_mort int,
	total_bc_limit int,
	total_il_high_credit_limit int,
	revol_bal_joint int,
	sec_app_earliest_cr_line int,
	sec_app_inq_last_6mths int,
	sec_app_mort_acc int,
	sec_app_open_acc int,
	sec_app_revol_util int,
	sec_app_open_act_il int,
	sec_app_num_rev_accts int,
	sec_app_chargeoff_within_12_mths int,
	sec_app_collections_12_mths_ex_med int,
	sec_app_mths_since_last_major_derog int,
	hardship_flag varchar(55),
	hardship_type varchar(55),
	hardship_reason varchar(55),
	hardship_status varchar(55),
	deferral_term varchar(55),
	hardship_amount int,
	hardship_start_date varchar(55),
	hardship_end_date varchar(55),
	payment_plan_start_date varchar(55),
	hardship_length int,
	hardship_dpd varchar(55),
	hardship_loan_status varchar(55),
	orig_projected_additional_accrued_interest int,
	hardship_payoff_balance_amount int,
	hardship_last_payment_amount int,
	disbursement_method varchar(55),
	debt_settlement_flag varchar(55),
	debt_settlement_flag_date varchar(55),
	settlement_status varchar(55),
	settlement_date varchar(55),
	settlement_amount int,
	settlement_percentage int,
	settlement_term varchar(55),
    CountryID int,
    Employer varchar(55),
	EmploymentType int,
	EmpOrderNum int,
	GrossMonthlyIncome int,
	LenEmpMons int,
	LenEmpYears int,
	Position varchar(55),
	RetiredFlag varchar(55),
	SelfEmpFlag varchar(55),
	`State` varchar(55),
    created_time TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
    updated_time TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (row_id)
);



create table applications (
    row_id varchar(55) not null,
    app_data json,
    employment_data json,
    created_time TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
    updated_time TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (row_id)
);
