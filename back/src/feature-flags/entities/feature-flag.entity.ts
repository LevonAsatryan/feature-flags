import { Company } from 'src/company/entities/company.entity';
import { Container } from 'src/containers/entities/container.entity';
import {
  Column,
  Entity,
  JoinColumn,
  ManyToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity()
export class FeatureFlag {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ type: 'varchar', length: 100 })
  name: string;

  @Column({ type: 'boolean' })
  isEnabled: boolean;

  @ManyToOne(() => Company, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'company_id' })
  companyId: number;

  @ManyToOne(() => Container, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'container_id' })
  containerId: number;
}
