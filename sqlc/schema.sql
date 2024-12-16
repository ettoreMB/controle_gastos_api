CREATE TABLE IF NOT EXISTS expenses (
  id uuid PRIMARY KEY DEFAULT(gen_random_uuid()),
  name varchar(100) NOT NULL, 
  value float NOT NULL,
  category_id uuid NOT NULL,
  payment_method_id uuid NOT NULL,
  status varchar(10) DEFAULT("pending"),
  credit_card_id UUID,
  installment_count int,
  current_installment int,
  installment_count_value float,
  transaction_date timestamptz NOT NULL DEFAULT now(),

  FOREIGN KEY (status) REFERENCES status(name)
);

CREATE TABLE  IF NOT EXISTS status  (
  name varchar(10) PRIMARY KEY DEFAULT("pending")
);

