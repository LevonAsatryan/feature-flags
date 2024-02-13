import { Company } from 'src/company/entities/company.entity';
import {
  Column,
  Entity,
  JoinColumn,
  ManyToMany,
  ManyToOne,
  OneToMany,
  OneToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity()
export class Container {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ type: 'varchar', length: 100 })
  name: string;

  @ManyToOne(() => Company, {
    onDelete: 'CASCADE',
  })
  @JoinColumn({ name: 'company_id' })
  companyId: number;

  @ManyToOne(() => Container)
  @JoinColumn({ name: 'parent_container_id' })
  parentContainer: Container;

  @OneToMany(() => Container, container => container.parentContainer)
  childrenContainers: Container[];
}
