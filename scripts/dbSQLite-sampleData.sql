BEGIN TRANSACTION;
INSERT INTO "Users" ("ID","Email","Password","Role","Name") VALUES (1,'joao@couve.com','5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5',1,'João das Couves'),
 (2,'lanchonete@dce.com','5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5',2,'Lanchonete DCE');
INSERT INTO "University" ("ID","Name") VALUES (1,'PUCPR - Câmpus Curitiba');
INSERT INTO "Restaurant" ("ID","University","Name") VALUES (1,1,'Lanchonete DCE');
INSERT INTO "UserReference" ("ID","University","Restaurant","User") VALUES (1,1,NULL,1),
 (2,NULL,1,2);
INSERT INTO "Item" ("ID","Restaurant","Price","Quantity","Name","Description","CancelGracePeriod") VALUES (1,1,4,50,'Coxinha','Uma unidade de coxinha',10);
INSERT INTO "Item" ("ID","Restaurant","Price","Quantity","Name","Description","CancelGracePeriod") VALUES (1,1,4,50,'Doguinho','Pão com salsicha',10);
COMMIT;